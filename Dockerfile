FROM --platform=$BUILDPLATFORM golang:1.20 AS builder

ARG VERSION
ARG COMMIT

ADD . $GOPATH/src/github.com/quantcdn/backend-init/

WORKDIR $GOPATH/src/github.com/quantcdn/backend-init

ENV CGO_ENABLED 0

ARG TARGETOS TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} && \
    go mod tidy && \
    go build -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" -o build/backend-init

FROM scratch

COPY --from=builder /go/src/github.com/quantcdn/backend-init/build/backend-init /usr/local/bin/backend-init

ENTRYPOINT [ "/usr/local/bin/backend-init" ]