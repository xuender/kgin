package kgin

import (
	"fmt"
	"net/http"
)

type NotFoundIDError uint64

func NewNotFoundIDError(err error, id uint64) error {
	if err == nil {
		return nil
	}

	return NotFoundIDError(id)
}

func (p NotFoundIDError) Error() string {
	return fmt.Sprintf("%d:%d", http.StatusNotFound, p)
}

func (p NotFoundIDError) String() string {
	return fmt.Sprintf("%d not found id", p)
}

type NotFoundKeyError string

func (p NotFoundKeyError) Error() string {
	return fmt.Sprintf("%d:%s", http.StatusNotFound, p.String())
}

func (p NotFoundKeyError) String() string {
	return fmt.Sprintf("%s not found key", string(p))
}

type NotFoundError string

func (p NotFoundError) Error() string {
	return fmt.Sprintf("%d:%s", http.StatusNotFound, string(p))
}

func (p NotFoundError) String() string {
	return string(p)
}
