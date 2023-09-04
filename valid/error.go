package valid

import (
	"fmt"
	"net/http"
)

type BadRequestError string

func NewBadRequestError(err error) error {
	if err == nil {
		return nil
	}

	return BadRequestError(err.Error())
}

func (p BadRequestError) Error() string {
	return fmt.Sprintf("%d:%s", http.StatusBadRequest, p.String())
}

func (p BadRequestError) String() string {
	return string(p)
}
