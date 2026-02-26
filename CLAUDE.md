# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a gRPC key-value store example with three components: a gRPC server, a gRPC-gateway (REST-to-gRPC proxy), and a CLI client.

## Build Commands

```bash
# Install protoc code generation tools
make install-deps

# Full build: regenerates protobufs then builds all three binaries
make build

# Build individual components
make server
make gateway
make client

# Regenerate protobuf stubs only
make protobufs         # generates kvstore.pb.go and kvstore_grpc.pb.go
make gateway-protobufs # generates kvstore.pb.gw.go

# Clean generated files and binaries (removes bin/* AND pkg/kvstore/*.go)
make clean

# Update all Go dependencies
make update-deps
```

The generated files in `pkg/kvstore/` (`*.pb.go`, `*.pb.gw.go`, `*_grpc.pb.go`) are derived from `pkg/kvstore/kvstore.proto` — do not edit them manually.

There are no tests in this repository.

## Running

```bash
# Start gRPC server (port 4000)
bin/server -port 4000

# Start REST gateway (port 8080, proxies to gRPC server)
bin/gateway -port 8080 -grpcServerAddr localhost:4000

# CLI client commands
bin/client -serverAddr localhost:4000 put <key> <value>
bin/client -serverAddr localhost:4000 get <key>
bin/client -serverAddr localhost:4000 -ssl put <key> <value>  # with TLS

# REST endpoints (requires gateway running)
curl -X PUT http://localhost:8080/v1 -d '{"key":"k","value":"v"}'
curl "http://localhost:8080/v1?key=k"

# Docker
make image     # builds Docker image (requires binaries already built via make build)
make run       # runs container exposing ports 4000 and 8080
make run-bash  # opens a shell in the container
```

## Architecture

```
cmd/
  server/   → entry point, delegates to pkg/server
  gateway/  → entry point, delegates to pkg/gateway
  client/   → entry point, CLI arg parsing, delegates to pkg/client
pkg/
  kvstore/  → protobuf definition (kvstore.proto) + generated stubs
              kvstore.yaml defines HTTP/REST routing rules for grpc-gateway
  server/   → gRPC server implementation (in-memory map store)
  gateway/  → grpc-gateway HTTP/REST reverse proxy
  client/   → gRPC client wrapper
  version/  → build version/branch injected via ldflags
```

**Data flow:**
- gRPC clients talk directly to `server` on port 4000
- REST clients talk to `gateway` on port 8080, which proxies to `server` on port 4000
- HTTP routing rules (GET /v1, PUT /v1) are configured in `pkg/kvstore/kvstore.yaml`, not annotations in the proto file

**Proto compilation** requires both `protoc` (system) and Go plugins (`protoc-gen-go`, `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway`). The two `make protobufs` targets must be run separately because they use different protoc invocations.

**Build output:** `make build` produces both a native binary (e.g. `bin/server`) and a cross-compiled `linux/amd64` binary (e.g. `bin/server-linux-amd64`) for each component. The linux binaries are what the Docker image uses.

**Notable implementation details:**
- The in-memory store (`pkg/server/server.go`) uses a plain `map[string]string` with no mutex — not safe for concurrent writes
- TLS mode on the client (`-ssl` flag) uses `InsecureSkipVerify: true` — suitable only for testing
- `pkg/version` exposes `BuildVersion`/`BuildBranch` vars injected via `-ldflags` at build time (only for the linux-amd64 targets)

## Dependencies

- `google.golang.org/grpc` — gRPC runtime
- `google.golang.org/protobuf` — protobuf runtime
- `github.com/grpc-ecosystem/grpc-gateway/v2` — REST-to-gRPC gateway
