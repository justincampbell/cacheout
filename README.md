# cacheout

> Caches output from commands for the specified timeframe

## Installation

1. Download the latest package for your platform from the [Releases page](https://github.com/justincampbell/cacheout/releases/latest).
2. Untar the package with `tar -zxvf cachout*.tar.gz`.
3. Move the extracted `cacheout` file to a directory in your `$PATH` (for most systems, this will be `/usr/local/bin/`).

Or, if you have a [Go development environment](https://golang.org/doc/install):

```
go get github.com/justincampbell/cacheout
```

## Usage

```
cacheout 1m command [args]
```

The duration is parsed with [Go's time.ParseDuration](https://golang.org/pkg/time/#ParseDuration).
Example durations are `30s`, `5m`, `2h`, or `1h30m`.
