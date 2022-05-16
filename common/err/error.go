package err

var (
	NilRouter = error{code: -1, mess: "router is nil"}

	NilRequest           = error{code: -2, mess: "request is nil"}
	TooLongCode          = error{code: -3, mess: "code is too long"}
	HashPasswordFailed   = error{code: -4, mess: "internal error"}
	AddMAccountFailed    = error{code: -5, mess: "internal error"}
	NilPassword          = error{code: -6, mess: "password is empty"}
	NilMerchantCode      = error{code: -7, mess: "merchant code is empty"}
	GetMAccountFailed    = error{code: -8, mess: "internal error"}
	UpdateMAccountFailed = error{code: -9, mess: "internal error"}
	DeleteMAccountFailed = error{code: -10, mess: "internal error"}
	AddMMemberFailed     = error{code: -11, mess: "internal error"}
	EmailExisted         = error{code: -12, mess: "email existed"}
	CheckExistenceFailed = error{code: -13, mess: "check existence failed"}
	MerchantCodeExisted  = error{code: -14, mess: "merchant code existed"}
	EmptyEmail           = error{code: -15, mess: "email is empty"}
	GetMMemberFailed     = error{code: -16, mess: "internal error"}
	UpdateMMemberFailed  = error{code: -17, mess: "internal error"}
)

type error struct {
	data interface{}
	code int
	mess string
}

func (e error) Data() interface{} {
	return e.data
}

func (e error) Code() int {
	return e.code
}

func (e error) Error() string {
	return e.mess
}

type Error struct {
	Domain  string `json:"domain"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
