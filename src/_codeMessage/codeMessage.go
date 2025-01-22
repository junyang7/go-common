package _codeMessage

type CodeMessage struct {
	Code    int
	Message string
}

func New(code int, message string) *CodeMessage {
	return &CodeMessage{
		Code:    code,
		Message: message,
	}
}

var (
	ErrNone    = New(0, "success")
	ErrDefault = New(1, "something goes wrong...!!!")
)
