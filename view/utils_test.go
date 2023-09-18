package view_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kgin/view"
)

func TestToUint64(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	ass.Equal(uint64(1), view.ToUint64([]byte{0, 0, 0, 0, 0, 0, 0, 1}))
	ass.Equal(uint64(1_000), view.ToUint64([]byte{0, 0, 0, 0, 0, 0, 3, 0xE8}))
}

func TestToBytes(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 1}, view.ToBytes(1))
	ass.Equal([]byte{0, 0, 0, 0, 0, 0, 3, 0xE8}, view.ToBytes(1_000))
}
