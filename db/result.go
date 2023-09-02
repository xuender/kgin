package db

import (
	"github.com/xuender/kvalid"
)

type Result[T kvalid.RuleHolder] struct {
	Count  int64 `json:"count,omitempty"`
	Limit  int   `json:"limit,omitempty"`
	Offset int   `json:"offset,omitempty"`
	Data   []T   `json:"data,omitempty"`
}
