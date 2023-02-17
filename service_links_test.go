package sutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Link(t *testing.T) {
	a := assert.New(t)

	waitFor := func(linkFunc, linkExpected int) func() {
		return func() {
			a.Equal(linkFunc, linkExpected)
		}
	}

	link := NewLink(true, true)
	link.Switch(nil, nil, nil, nil)

	link = NewLink(true, true)
	link.Switch(
		waitFor(linkIxI, linkIxI),
		waitFor(linkExI, linkIxI),
		waitFor(linkIxE, linkIxI),
		waitFor(linkExE, linkIxI),
	)

	link = NewLink(false, true)
	link.Switch(
		waitFor(linkIxI, linkExI),
		waitFor(linkExI, linkExI),
		waitFor(linkIxE, linkExI),
		waitFor(linkExE, linkExI),
	)

	link = NewLink(true, false)
	link.Switch(
		waitFor(linkIxI, linkIxE),
		waitFor(linkExI, linkIxE),
		waitFor(linkIxE, linkIxE),
		waitFor(linkExE, linkIxE),
	)

	link = NewLink(false, false)
	link.Switch(
		waitFor(linkIxI, linkExE),
		waitFor(linkExI, linkExE),
		waitFor(linkIxE, linkExE),
		waitFor(linkExE, linkExE),
	)
}
