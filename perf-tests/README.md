# Performance tests

Load tests using [resonate](https://github.com/jecklgamis/resonate).

```sh
brew install jecklgamis/tap/resonate
```

## Key-value store (`kvstore.KeyValueStore`)

Start the server on port 4000, then run:

```sh
go run cmd/server/server.go -port 4000
./run-perf-tests.sh
```

Each iteration is a `Put` followed by a `Get` of the same key, ramping to 100 iterations/s
(about 200 requests/s) with 50 workers, holding for 1 minute, then ramping down. It fails
(non-zero exit) if success rate is below 99% or p95 latency exceeds 200ms.
Change `grpc.address` in the scenario to target another host.
