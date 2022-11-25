package _exception

type Exception struct {
	Code    int
	Message string
}

func New(code int, message string) *Exception {
	return &Exception{
		Code:    code,
		Message: message,
	}
}
