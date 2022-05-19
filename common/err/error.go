package err

import "main.go/common/status"

var (
	NilRouter = InternalError{Code: -1, Mess: "router is nil"}

	NilRequest           = InternalError{Code: -2, Mess: "request is nil"}
	TooLongCode          = InternalError{Code: -3, Mess: "Code is too long"}
	HashPasswordFailed   = InternalError{Code: -4, Mess: "internal error"}
	AddMAccountFailed    = InternalError{Code: -5, Mess: "internal error"}
	EmptyPassword        = InternalError{Code: -6, Mess: "password is empty"}
	EmptyMerchantCode    = InternalError{Code: -7, Mess: "merchant Code is empty"}
	GetMAccountFailed    = InternalError{Code: -8, Mess: "internal error"}
	UpdateMAccountFailed = InternalError{Code: -9, Mess: "internal error"}
	DeleteMAccountFailed = InternalError{Code: -10, Mess: "internal error"}
	AddMMemberFailed     = InternalError{Code: -11, Mess: "internal error"}
	EmailExisted         = InternalError{Code: -12, Mess: "email existed"}
	CheckExistenceFailed = InternalError{Code: -13, Mess: "check existence failed"}
	MerchantCodeExisted  = InternalError{Code: -14, Mess: "merchant Code existed"}
	EmptyEmail           = InternalError{Code: -15, Mess: "email is empty"}
	GetMMemberFailed     = InternalError{Code: -16, Mess: "internal error"}
	UpdateMMemberFailed  = InternalError{Code: -17, Mess: "internal error"}
	DeleteMMemberFailed  = InternalError{Code: -18, Mess: "internal error"}

	EmptyMerchantName = InternalError{Code: -19, Mess: "merchant name is empty"}
	EmptyUserName     = InternalError{Code: -20, Mess: "username is empty"}
	NotFound          = InternalError{Code: -21, Mess: "item not found"}
	InvalidParameter  = InternalError{Code: -22, Mess: "parameter is invalid"}
	EmptyMerchantID    = InternalError{Code: -23, Mess: "merchant id is empty"}
)

type InternalError struct {
	Data interface{}
	Code int
	Mess string
}

func (e InternalError) Error() string {
	return e.Mess
}

func (e InternalError) ToExternalError(er error) *Error {
	if customErr, ok := er.(InternalError); ok {
		return &Error{
			Domain:  status.Domain,
			Code:    customErr.Code,
			Message: customErr.Error(),
		}
	}
	return &Error{
		Domain:  status.Domain,
		Code:    e.Code,
		Message: e.Error(),
	}
}

type Error struct {
	Domain  string `json:"domain"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
