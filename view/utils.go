package view

import (
	"encoding/binary"

	"github.com/xuender/kit/times"
	"github.com/xuender/kit/types"
)

// nolint: gochecknoglobals
var (
	_pv    = [...]byte{'p', 'v', '_'}
	_uv    = [...]byte{'u', 'v', '_'}
	_count = [...]byte{'c', 'o', '_'}
)

func CountKey(page string) []byte {
	var (
		key = pageKey(page)
		ret = make([]byte, 0, len(key)+len(_count))
	)
	// co_ + key
	ret = append(ret, _count[:]...)
	ret = append(ret, key...)

	return ret
}

func pageKey(page string) []byte {
	if pid, err := types.ParseInteger[uint64](page); err == nil {
		return ToBytes(pid)
	}

	return []byte(page)
}

func PVKey(page string, day times.IntDay) []byte {
	var (
		length = 7
		key    = pageKey(page)
		ret    = make([]byte, 0, len(key)+length)
	)
	// pv_ + key + 日期
	ret = append(ret, _pv[:]...)
	ret = append(ret, key...)
	ret = append(ret, day.Marshal()...)

	return ret
}

func UVKey(page string, day times.IntDay) []byte {
	var (
		length = 7
		key    = pageKey(page)
		ret    = make([]byte, 0, len(key)+length)
	)
	// uv_ + 日期 + key
	ret = append(ret, _uv[:]...)
	ret = append(ret, day.Marshal()...)
	ret = append(ret, key...)

	return ret
}

func ToUint64(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}

func ToBytes(count uint64) []byte {
	length := 8
	data := make([]byte, length)
	binary.BigEndian.PutUint64(data, count)

	return data
}
