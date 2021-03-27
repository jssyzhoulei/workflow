package code

type Code int32

func (c Code) Message(err ...error) string {
	if m, ok := message[c]; ok {
		return m
	}
	if err != nil {
		return err[0].Error()
	}
	return "未知的错误"
}
const OK Code = 200
const (
	PARAMS_ERROR Code = iota + 200100
	XlsxError
)

const (
	SVC_ERROR Code = iota + 200200
)
