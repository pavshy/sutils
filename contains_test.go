package sutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntInSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	assert.True(t, IntInSlice(1, intSlice))
	assert.False(t, IntInSlice(6, intSlice))
}

func TestStringInSlice(t *testing.T) {
	stringSlice := []string{"hello", "world", "here", "we", "go", "again"}

	assert.True(t, StringInSlice("hello", stringSlice))
	assert.False(t, StringInSlice("Hello", stringSlice))
	assert.False(t, StringInSlice("notExist", stringSlice))
}
