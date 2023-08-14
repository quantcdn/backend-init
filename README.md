# Backend initialisation

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/quantcdn/backend-init)
[![Go Report Card](https://goreportcard.com/badge/github.com/quantcdn/backend-init)](https://goreportcard.com/report/github.com/quantcdn/backend-init)
[![Coverage Status](https://coveralls.io/repos/github/quantcdn/backend-init/badge.svg?branch=main)](https://coveralls.io/github/quantcdn/backend-init?branch=main)
[![Release](https://img.shields.io/github/v/release/quantcdn/backend-init)](https://github.com/quantcdn/backend-init/releases/latest)

This is intended to be run as an [init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/). It will make HTTP requests to a configured URL to determine if the service is functional.

## Installation

As this is intended to be a docker entrypoint the preferred way to install is using with a dockerfile.

```Dockerfile
COPY --from=ghcr.io/quantcdn/backend-init:latest /usr/local/bin/backend-init /usr/local/bin/backend-init
```

This can be run directly from the docker image:

```sh
docker run --rm ghcr.io/quantcdn/backend-init:latest backend-init --version
```

## Usage

```
$ backend-init --help                                                                                                                             
usage: backend-init [<flags>]

Flags:
  --help      Show context-sensitive help (also try --help-long and --help-man).
  --url=URL   The backend url.
  --retry=10  Times to retry the backend connection.
  --delay=5   Delay between backend requests.
  --version   Show application version.
```

For example to execute `build` and `start` after a backend connection has been established.

```
backend-init --url http://localhost build start
```

## Local development

### Build
```sh
git clone git@github.com:quantcdn/backend-init.git && cd backend-init
go generate ./...
go build -ldflags="-s -w" -o build/backend-init .
go run . -h
```

### Run tests
```sh
go test -v ./... -coverprofile=build/coverage.out
```

View coverage results:
```sh
go tool cover -html=build/coverage.out
```

### Documentation
```sh
cd docs
npm install
npm run dev
```