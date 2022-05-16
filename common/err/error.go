package err

var (
	NilRouter = error{data: "router is nil", code: -1}
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
	Domain string `json:"domain"`
	Code int `json:"code"`
	Message string `json:"message"`
}