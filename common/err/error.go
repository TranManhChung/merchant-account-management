package err

var (
	NilRouter = InternalError{code: -1, mess: "router is nil"}

	NilRequest           = InternalError{code: -2, mess: "request is nil"}
	TooLongCode          = InternalError{code: -3, mess: "code is too long"}
	HashPasswordFailed   = InternalError{code: -4, mess: "internal error"}
	AddMAccountFailed    = InternalError{code: -5, mess: "internal error"}
	EmptyPassword        = InternalError{code: -6, mess: "password is empty"}
	EmptyMerchantCode    = InternalError{code: -7, mess: "merchant code is empty"}
	GetMAccountFailed    = InternalError{code: -8, mess: "internal error"}
	UpdateMAccountFailed = InternalError{code: -9, mess: "internal error"}
	DeleteMAccountFailed = InternalError{code: -10, mess: "internal error"}
	AddMMemberFailed     = InternalError{code: -11, mess: "internal error"}
	EmailExisted         = InternalError{code: -12, mess: "email existed"}
	CheckExistenceFailed = InternalError{code: -13, mess: "check existence failed"}
	MerchantCodeExisted  = InternalError{code: -14, mess: "merchant code existed"}
	EmptyEmail           = InternalError{code: -15, mess: "email is empty"}
	GetMMemberFailed     = InternalError{code: -16, mess: "internal error"}
	UpdateMMemberFailed  = InternalError{code: -17, mess: "internal error"}
	DeleteMMemberFailed  = InternalError{code: -18, mess: "internal error"}

	EmptyMerchantName = InternalError{code: -19, mess: "merchant name is empty"}
	EmptyUserName     = InternalError{code: -20, mess: "username is empty"}
	NotFound          = InternalError{code: -21, mess: "item not found"}
)

type InternalError struct {
	data interface{}
	code int
	mess string
}

func (e InternalError) Data() interface{} {
	return e.data
}

func (e InternalError) Code() int {
	return e.code
}

func (e InternalError) Error() string {
	return e.mess
}

type Error struct {
	Domain  string `json:"domain"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
