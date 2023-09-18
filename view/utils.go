package view

import (
	"encoding/binary"

	"github.com/xuender/kit/times"
)

// nolint: gochecknoglobals
var (
	_pv    = [...]byte{'p', 'v', '_'}
	_uv    = [...]byte{'u', 'v', '_'}
	_count = [...]byte{'c', 'o', '_'}
)

func CountKey(key uint64) []byte {
	var (
		length   = 11
		keyBytes = ToBytes(key)
		ret      = make([]byte, 0, length)
	)
	// co_ + key
	ret = append(ret, _count[:]...)
	ret = append(ret, keyBytes...)

	return ret
}

func PVKey(key uint64, day times.IntDay) []byte {
	var (
		length   = 15
		keyBytes = ToBytes(key)
		ret      = make([]byte, 0, length)
	)
	// pv_ + key + 日期
	ret = append(ret, _pv[:]...)
	ret = append(ret, keyBytes...)
	ret = append(ret, day.Marshal()...)

	return ret
}

func UVKey(key uint64, day times.IntDay) []byte {
	var (
		length   = 15
		keyBytes = ToBytes(key)
		ret      = make([]byte, 0, length)
	)
	// uv_ + 日期 + key
	ret = append(ret, _uv[:]...)
	ret = append(ret, day.Marshal()...)
	ret = append(ret, keyBytes...)

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
