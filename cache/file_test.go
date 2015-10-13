package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FileCache(t *testing.T) {
	key := "abc123"
	fc := NewFileCache(key)

	// Check that FileCache implements the Cache interface
	var _ Cache = fc

	// Write
	bytes := []byte("foobarbaz")
	_, err := fc.Write(bytes)
	assert.Nil(t, err)

	// Persist
	err = fc.Persist()
	assert.Nil(t, err)

	// Read
	assert.Equal(t, bytes, fc.Bytes())

	// Read from another cache with the same key
	fc = NewFileCache(key)
	assert.Equal(t, bytes, fc.Bytes())

	// A new cache should be empty
	assert.Equal(t, []byte{}, NewFileCache("anotherkey").Bytes())
}
