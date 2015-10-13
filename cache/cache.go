// Package cache provide caching implementations.
package cache

// Cache is the interface that each cache implementation must implement.
type Cache interface {
	Write([]byte) (int, error)
	Persist() error
	Bytes() []byte
}
