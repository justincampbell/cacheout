package cache

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// FileCache is a Cache implementation which caches to the filesystem.
type FileCache struct {
	Key string

	buffer    *bytes.Buffer
	cachePath string
}

const (
	fileMode = 0644
	prefix   = "cacheout"
)

// NewFileCache returns a FileCache with an empty buffer and precomputed cache
// path based on the key.
func NewFileCache(key string) *FileCache {
	return &FileCache{
		Key:       key,
		buffer:    bytes.NewBuffer([]byte{}),
		cachePath: path.Join(os.TempDir(), fmt.Sprintf("%s.%s", prefix, key)),
	}
}

// Write writes bytes to the internal buffer.
func (fc *FileCache) Write(bytes []byte) (int, error) {
	return fc.buffer.Write(bytes)
}

// Persist stores the internal buffer to the cache file. The file is emptied if
// it already exists.
func (fc *FileCache) Persist() error {
	return ioutil.WriteFile(
		fc.cachePath,
		fc.buffer.Bytes(),
		fileMode,
	)
}

// Bytes returns the bytes from the cache file, or an empty byte array if the
// cache file does not exist.
func (fc *FileCache) Bytes() []byte {
	bytes, err := ioutil.ReadFile(fc.cachePath)
	if err != nil {
		return []byte{}
	}

	return bytes
}

// Stale returns whether or not the cache is stale for the given TTL.
func (fc *FileCache) Stale() bool {
	_, err := os.Stat(fc.cachePath)

	return os.IsNotExist(err)
}

// Expire clears the cache.
func (fc *FileCache) Expire() {
	_ = os.Remove(fc.cachePath)
}
