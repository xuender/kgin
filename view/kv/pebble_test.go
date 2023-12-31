package kv_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kgin/view/kv"
	"github.com/xuender/kit/oss"
	"github.com/xuender/kit/times"
)

// nolint: paralleltest
func TestPebble_PV(t *testing.T) {
	ass := assert.New(t)
	req := require.New(t)
	path := filepath.Join(os.TempDir(), "test_pv")
	peb, err := kv.NewPebble(path)
	page := uint64(23)
	day := times.Now2IntDay()

	req.NoError(err)

	ass.Equal(uint64(0), peb.PV(page, day))
	ass.Equal(uint64(0), peb.UV(page, day))
	req.NoError(peb.View(page, "127.0.0.1"))
	ass.Equal(uint64(1), peb.PV(page, day))
	ass.Equal(uint64(1), peb.UV(page, day))
	req.NoError(peb.View(page, "127.0.0.1"))
	ass.Equal(uint64(2), peb.PV(page, day))
	ass.Equal(uint64(1), peb.UV(page, day))
	req.NoError(peb.View(page, "127.0.0.2"))
	ass.Equal(uint64(3), peb.PV(page, day))
	ass.Equal(uint64(2), peb.UV(page, day))
	ass.Equal(uint64(3), peb.TV(page))

	day--

	ass.Equal(uint64(0), peb.PV(page, day))
	ass.Equal(uint64(0), peb.UV(page, day))

	patches := gomonkey.ApplyFuncReturn(times.Now2IntDay, day)

	req.NoError(peb.View(page, "127.0.0.1"))
	patches.Reset()
	ass.Equal(uint64(1), peb.PV(page, day))
	ass.Equal(uint64(1), peb.UV(page, day))
	ass.Equal(uint64(4), peb.TV(page))

	peb.Close()

	_ = oss.Remove(path, 0)
}
