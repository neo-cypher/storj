ARG DOCKER_ARCH
ARG DOCKER_PLATFORM

# TODO: specify an exact version. Should we do it on release branch?
FROM --platform=${DOCKER_PLATFORM:-linux/amd64} storjlabs/storagenode-base:latest-${DOCKER_ARCH:-amd64}
ARG TAG
ARG GOARCH
ARG VERSION_SERVER_URL
ENV GOARCH ${GOARCH:-amd64}
ENV VERSION_SERVER_URL ${VERSION_SERVER_URL:-https://version.storj.io}
EXPOSE 28967
EXPOSE 14002
COPY cmd/storagenode/docker/ /
WORKDIR /app
ENTRYPOINT ["/entrypoint"]

ENV ADDRESS="" \
    EMAIL="" \
    WALLET="" \
    STORAGE="2.0TB" \
    SETUP="false" \
    AUTO_UPDATE="true"
