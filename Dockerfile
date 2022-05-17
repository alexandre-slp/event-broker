FROM golang:1.17-alpine AS base

RUN apk add curl build-base
WORKDIR /event-broker


FROM base as debug

RUN go install github.com/cortesi/modd/cmd/modd@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin v1.43.0

CMD ['go', 'run', '-race', './cmd/server/main.go']


FROM base as dev

CMD ['go', 'run', '-race', './cmd/server/main.go']


FROM base AS compiler_server
ARG SERVER_BIN_PATH
ARG SERVER_CMD_PATH

COPY ./ ./
RUN go mod vendor
RUN go build -o /bin/server ./cmd/server


FROM alpine AS release
ARG RELEASE_BIN_ARGS
ARG RELEASE_BIN_PATH

ENV RELEASE_BIN_ARGS ${RELEASE_BIN_ARGS}
ENV RELEASE_BIN_PATH ${RELEASE_BIN_PATH}
ENV TINI_VERSION v0.19.0

RUN apk add --update --no-cache ca-certificates tzdata && \
    ln -fs /usr/share/zoneinfo/UTC /etc/localtime && \
    rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

COPY --from=compiler_server /bin/server /bin/server

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /bin/tini
RUN chmod +x /bin/tini

RUN addgroup -g 1000 -S nonroot && \
    adduser -u 1000 -S nonroot -G nonroot
USER nonroot

ENTRYPOINT ['/bin/tini', '--']

CMD ${RELEASE_BIN_PATH} ${RELEASE_BIN_ARGS}