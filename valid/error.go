package valid

import (
	"fmt"
	"net/http"
)

type BadRequestError struct {
	err error
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{err}
}

func (p BadRequestError) Error() string {
	return fmt.Sprintf("%d:%s", http.StatusBadRequest, p.err)
}

func (p BadRequestError) Err() string {
	return p.err.Error()
}
