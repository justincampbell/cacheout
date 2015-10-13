// Package cache provide caching implementations.
package cache

// Cache is the interface that each cache implementation must implement.
type Cache interface {
	Bytes() []byte
	Expire()
	Persist() error
	Stale() bool
	Write([]byte) (int, error)
}
