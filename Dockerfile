FROM golang:1.17-alpine AS base

WORKDIR /event-broker

# Installing necessary components
RUN apk add curl build-base


FROM base as development
RUN go install github.com/cortesi/modd/cmd/modd@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.0
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

CMD [ "modd", "-f", "./cmd/server/modd.conf" ]


FROM base AS compiler
COPY . ./
RUN go mod vendor


FROM compiler AS compiler_server
ARG SERVER_BIN_PATH
ARG SERVER_CMD_PATH
RUN go build -o /bin/server ./cmd/server


FROM alpine AS release
ARG BIN_PATH
ARG BIN_ARGS

ENV BIN_ARGS ${BIN_ARGS}
ENV BIN_PATH ${BIN_PATH}

RUN apk add --update --no-cache ca-certificates tzdata \
    && ln -fs /usr/share/zoneinfo/UTC /etc/localtime \
    && rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

COPY --from=compiler_server /bin/server /bin/server

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /bin/tini
RUN chmod +x /bin/tini

RUN addgroup -g 1000 -S nonroot && \
    adduser -u 1000 -S nonroot -G nonroot
USER nonroot

ENTRYPOINT ["/bin/tini", "--"]

CMD ${BIN_PATH} ${BIN_ARGS}