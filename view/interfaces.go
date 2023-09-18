package view

import "github.com/xuender/kit/times"

type Viewer interface {
	View(page string, remoteIP string)
}

type Stat interface {
	PV(page string, day times.IntDay) uint64
	UV(page string, day times.IntDay) uint64
	Count(page string) uint64
}
