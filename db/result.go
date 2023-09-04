package db

type Result[T any] struct {
	Count  int64 `json:"count,omitempty"`
	Limit  int   `json:"limit,omitempty"`
	Offset int   `json:"offset,omitempty"`
	Data   []T   `json:"data,omitempty"`
}
