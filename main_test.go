package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseArgs(t *testing.T) {
	command, err := parseArgs([]string{"1m", "echo", "hello"})
	assert.Nil(t, err)
	assert.Equal(t, "1m0s", command.ttl.String())
	assert.Equal(t, "echo", command.bin)
	assert.Equal(t, []string{"hello"}, command.args)
}

func Test_parseArgs_not_enough(t *testing.T) {
	_, err := parseArgs([]string{})
	assert.EqualError(t, err, "not enough arguments")

	_, err = parseArgs([]string{"1m"})
	assert.EqualError(t, err, "not enough arguments")
}

func Test_parseArgs_invalid_duration(t *testing.T) {
	command, err := parseArgs([]string{"0", "echo"})
	assert.Nil(t, err)
	assert.Equal(t, "0", command.ttl.String())

	_, err = parseArgs([]string{"1", "echo"})
	assert.Contains(t, err.Error(), "missing unit")

	_, err = parseArgs([]string{"foobar", "echo"})
	assert.Contains(t, err.Error(), "invalid duration")
}

func Test_hashCommand(t *testing.T) {
	a := []string{"foo"}
	b := []string{"bar", "baz"}
	assert.Equal(t, hashCommand(a), hashCommand(a))
	assert.NotEqual(t, hashCommand(a), hashCommand(b))
}
