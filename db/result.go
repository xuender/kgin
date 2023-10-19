package db

type Result[T any] struct {
	Count  int64 `json:"count"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Data   []T   `json:"data"`
}
