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

	// Start from a clean state
	fc.Expire()

	// Unused cache is stale
	assert.True(t, fc.Stale())

	// Write
	bytes := []byte("foobarbaz")
	_, err := fc.Write(bytes)
	assert.Nil(t, err)

	// Persist
	err = fc.Persist()
	assert.Nil(t, err)

	// Freshly persisted cache is not stale
	assert.False(t, fc.Stale())

	// Read
	assert.Equal(t, bytes, fc.Bytes())

	// Read from another cache with the same key
	fc = NewFileCache(key)
	assert.Equal(t, bytes, fc.Bytes())

	// Expire
	fc.Expire()
	assert.True(t, fc.Stale())
	assert.Equal(t, []byte{}, fc.Bytes())

	// A new cache should be empty
	assert.Equal(t, []byte{}, NewFileCache("anotherkey").Bytes())
}
