package view

import (
	"github.com/xuender/kit/times"
)

type Viewer interface {
	View(key uint64, remoteIP string) error
}

type Stat interface {
	PV(key uint64, day times.IntDay) uint64
	UV(key uint64, day times.IntDay) uint64
	TV(key uint64) uint64
}

type Counter interface {
	Viewer
	Stat
}
