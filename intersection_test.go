package sutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlicesIntersect(t *testing.T) {
	stringSlice1 := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit"}
	stringSlice2 := []string{"here", "we", "have", "only", "amet"}

	assert.True(t, StringSlicesIntersect(stringSlice1, stringSlice2))
	assert.False(t, StringSlicesIntersect(stringSlice1, []string{"here", "we", "have"}))
}
