// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package apigen_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"storj.io/common/testcontext"
	"storj.io/common/uuid"
	"storj.io/storj/private/api"
	"storj.io/storj/private/apigen"
	"storj.io/storj/private/apigen/example"
)

type (
	auth     struct{}
	service  struct{}
	response = struct {
		ID        uuid.UUID
		Date      time.Time
		PathParam string
		Body      string
	}
)

func (a auth) IsAuthenticated(ctx context.Context, r *http.Request, isCookieAuth, isKeyAuth bool) (context.Context, error) {
	return ctx, nil
}

func (a auth) RemoveAuthCookie(w http.ResponseWriter) {}

func (s service) GenTestAPI(ctx context.Context, pathParam string, id uuid.UUID, date time.Time, body struct{ Content string }) (*response, api.HTTPError) {
	return &response{
		ID:        id,
		Date:      date,
		PathParam: pathParam,
		Body:      body.Content,
	}, api.HTTPError{}
}

func send(ctx context.Context, method string, url string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader = http.NoBody
	if body != nil {
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(bodyJSON)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return respBody, nil
}

func TestAPIServer(t *testing.T) {
	ctx := testcontext.NewWithTimeout(t, 5*time.Second)
	defer ctx.Cleanup()

	router := mux.NewRouter()
	example.NewTestAPI(zaptest.NewLogger(t), monkit.Package(), service{}, router, auth{})

	server := httptest.NewServer(router)
	defer server.Close()

	id, err := uuid.New()
	require.NoError(t, err)

	expected := response{
		ID:        id,
		Date:      time.Now(),
		PathParam: "foo",
		Body:      "bar",
	}

	resp, err := send(ctx, http.MethodPost,
		fmt.Sprintf("%s/api/v0/testapi/%s?id=%s&date=%s",
			server.URL,
			expected.PathParam,
			url.QueryEscape(expected.ID.String()),
			url.QueryEscape(expected.Date.Format(apigen.DateFormat)),
		), struct{ Content string }{expected.Body},
	)
	require.NoError(t, err)

	var actual map[string]string
	require.NoError(t, json.Unmarshal(resp, &actual))

	for _, key := range []string{"ID", "Date", "PathParam", "Body"} {
		require.Contains(t, actual, key)
	}
	require.Equal(t, expected.ID.String(), actual["ID"])
	require.Equal(t, expected.Date.Format(apigen.DateFormat), actual["Date"])
	require.Equal(t, expected.Body, actual["Body"])
}
