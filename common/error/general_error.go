package error

import (
	"tchen/jwt-auth/common/enums"
)

type GeneralError struct {
	enums.ResponseBase
	ErrorTrace string `json:"error_trace"`
}

func (g GeneralError) Error() string {
	return g.ErrorTrace
}
