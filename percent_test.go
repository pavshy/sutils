package sutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Percent_String(t *testing.T) {
	a := assert.New(t)
	p := Percent(0.34)
	str := fmt.Sprintf("%v", p)
	a.Equal("34%", str)
}
