// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"storj.io/common/fpath"
	"storj.io/common/peertls/tlsopts"
	"storj.io/common/rpc"
	"storj.io/common/storj"
	"storj.io/private/cfgstruct"
	"storj.io/private/process"
	"storj.io/storj/multinode"
	"storj.io/storj/multinode/multinodedb"
	"storj.io/storj/multinode/nodes"
	"storj.io/storj/private/multinodeauth"
)

// Config defines multinode configuration.
type Config struct {
	Database string `help:"multinode database connection string" default:"sqlite3://file:$CONFDIR/master.db"`

	multinode.Config
}

var (
	rootCmd = &cobra.Command{
		Use:   "multinode",
		Short: "Multinode Dashboard",
	}
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the multinode dashboard",
		RunE:  cmdRun,
	}
	setupCmd = &cobra.Command{
		Use:         "setup",
		Short:       "Create config files",
		RunE:        cmdSetup,
		Annotations: map[string]string{"type": "setup"},
	}

	addCmd = &cobra.Command{
		Use:   "add [file]",
		Short: "Add storage node(s) from file or stdin to multinode dashboard",
		RunE:  cmdAdd,
		Args:  cobra.MaximumNArgs(1),
		Example: `
# add nodes from json file containing array of nodes data
$ multinode add nodes.json

# add node from json file containing a single node object
$ multinode add node.json

# read nodes data from stdin
$ cat nodes.json | multinode add -
`,
	}

	runCfg   Config
	setupCfg Config
	addCfg   struct {
		NodeID        string `help:"ID of the storage node" default:""`
		Name          string `help:"Name of the storage node" default:""`
		APISecret     string `help:"API Secret of the storage node" default:""`
		PublicAddress string `help:"Public IP Address of the storage node" default:""`

		Config
	}
	confDir     string
	identityDir string
)

func main() {
	process.ExecCustomDebug(rootCmd)
}

func init() {
	defaultConfDir := fpath.ApplicationDir("storj", "multinode")
	defaultIdentityDir := fpath.ApplicationDir("storj", "identity", "multinode")
	cfgstruct.SetupFlag(zap.L(), rootCmd, &confDir, "config-dir", defaultConfDir, "main directory for multinode configuration")
	cfgstruct.SetupFlag(zap.L(), rootCmd, &identityDir, "identity-dir", defaultIdentityDir, "main directory for multinode identity credentials")
	defaults := cfgstruct.DefaultsFlag(rootCmd)

	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(addCmd)

	process.Bind(runCmd, &runCfg, defaults, cfgstruct.ConfDir(confDir), cfgstruct.IdentityDir(identityDir))
	process.Bind(setupCmd, &setupCfg, defaults, cfgstruct.ConfDir(confDir), cfgstruct.IdentityDir(identityDir), cfgstruct.SetupMode())
	process.Bind(addCmd, &addCfg, defaults, cfgstruct.ConfDir(confDir), cfgstruct.IdentityDir(identityDir))
}

func cmdSetup(cmd *cobra.Command, args []string) (err error) {
	setupDir, err := filepath.Abs(confDir)
	if err != nil {
		return err
	}

	valid, _ := fpath.IsValidSetupDir(setupDir)
	if !valid {
		return fmt.Errorf("multinode configuration already exists (%v)", setupDir)
	}

	err = os.MkdirAll(setupDir, 0700)
	if err != nil {
		return err
	}

	return process.SaveConfig(cmd, filepath.Join(setupDir, "config.yaml"))
}

func cmdRun(cmd *cobra.Command, args []string) (err error) {
	ctx, _ := process.Ctx(cmd)
	log := zap.L()

	runCfg.Debug.Address = *process.DebugAddrFlag

	identity, err := runCfg.Identity.Load()
	if err != nil {
		log.Error("failed to load identity", zap.Error(err))
		return errs.New("failed to load identity: %+v", err)
	}

	db, err := multinodedb.Open(ctx, log.Named("db"), runCfg.Database)
	if err != nil {
		return errs.New("error connecting to master database on multinode: %+v", err)
	}
	defer func() {
		err = errs.Combine(err, db.Close())
	}()
	if err := db.MigrateToLatest(ctx); err != nil {
		return err
	}

	peer, err := multinode.New(log, identity, runCfg.Config, db)
	if err != nil {
		return err
	}

	runError := peer.Run(ctx)
	closeError := peer.Close()
	return errs.Combine(runError, closeError)
}

func cmdAdd(cmd *cobra.Command, args []string) (err error) {
	ctx, _ := process.Ctx(cmd)
	log := zap.L()

	identity, err := addCfg.Identity.Load()
	if err != nil {
		return errs.New("failed to load identity: %+v", err)
	}

	db, err := multinodedb.Open(ctx, log.Named("db"), addCfg.Database)
	if err != nil {
		return errs.New("error connecting to master database on multinode: %+v", err)
	}

	tlsConfig := tlsopts.Config{
		UsePeerCAWhitelist: false,
		PeerIDVersions:     "0",
	}

	tlsOptions, err := tlsopts.NewOptions(identity, tlsConfig, nil)
	if err != nil {
		return err
	}

	dialer := rpc.NewDefaultDialer(tlsOptions)

	var nodeList []nodes.Node

	hasRequiredFlags := addCfg.NodeID != "" && addCfg.APISecret != "" && addCfg.PublicAddress != ""

	if len(args) == 0 && !hasRequiredFlags {
		return errs.New("--node-id, --api-secret and --public-address flags are required if no file is provided")
	}

	if hasRequiredFlags {
		nodeID, err := storj.NodeIDFromString(addCfg.NodeID)
		if err != nil {
			return err
		}
		apiSecret, err := multinodeauth.SecretFromBase64(addCfg.APISecret)
		if err != nil {
			return err
		}
		nodeList = []nodes.Node{
			{
				ID:            nodeID,
				PublicAddress: addCfg.PublicAddress,
				APISecret:     apiSecret,
				Name:          addCfg.Name,
			},
		}
	} else {
		path := args[0]
		var nodesJSONData []byte
		if path == "-" {
			stdin := cmd.InOrStdin()
			data, err := io.ReadAll(stdin)
			if err != nil {
				return err
			}
			nodesJSONData = data
		} else {
			nodesJSONData, err = os.ReadFile(path)
			if err != nil {
				return err
			}
		}

		nodeList, err = unmarshalJSONNodes(nodesJSONData)
		if err != nil {
			return err
		}
	}

	for _, node := range nodeList {
		if _, err := db.Nodes().Get(ctx, node.ID); err == nil {
			return errs.New("Node with ID %s is already added to the multinode dashboard", node.ID)
		}

		service := nodes.NewService(log, dialer, db.Nodes())
		err = service.Add(ctx, node)
		if err != nil {
			return err
		}
	}

	return nil
}

// decodeUTF16or8 decodes the b as UTF-16 if the special byte order mark is present.
func decodeUTF16or8(b []byte) ([]byte, error) {
	r := bytes.NewReader(b)
	// fallback to r if no BOM sequence is located in the source text.
	t := unicode.BOMOverride(transform.Nop)
	return io.ReadAll(transform.NewReader(r, t))
}

func unmarshalJSONNodes(nodesJSONData []byte) ([]nodes.Node, error) {
	var nodesInfo []nodes.Node
	var err error

	nodesJSONData, err = decodeUTF16or8(nodesJSONData)
	if err != nil {
		return nil, err
	}
	nodesJSONData = bytes.TrimLeft(nodesJSONData, " \t\r\n")

	switch {
	case len(nodesJSONData) > 0 && nodesJSONData[0] == '[': // data is json array
		err := json.Unmarshal(nodesJSONData, &nodesInfo)
		if err != nil {
			return nil, err
		}
	case len(nodesJSONData) > 0 && nodesJSONData[0] == '{': // data is json object
		var singleNode nodes.Node
		err := json.Unmarshal(nodesJSONData, &singleNode)
		if err != nil {
			return nil, err
		}
		nodesInfo = []nodes.Node{singleNode}
	default:
		return nil, errs.New("invalid JSON format")
	}

	return nodesInfo, nil
}
