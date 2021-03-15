package apierror

import "fmt"

type ApiError struct {
	Err error
	Code int
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Err.Error())
}