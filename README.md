# cacheout

> Caches output from commands for the specified timeframe

## Installation

```
go get github.com/justincampbell/cacheout
```

## Usage

```
cacheout 1m command [args]
```

The duration is parsed with [Go's time.ParseDuration](https://golang.org/pkg/time/#ParseDuration).
Example durations are `30s`, `5m`, `2h`, or `1h30m`.
