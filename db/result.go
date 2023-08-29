package db

import "github.com/xuender/kgin/valid"

type Result[T valid.Valid] struct {
	Count  int64 `json:"count,omitempty"`
	Limit  int   `json:"limit,omitempty"`
	Offset int   `json:"offset,omitempty"`
	Data   []T   `json:"data,omitempty"`
}
