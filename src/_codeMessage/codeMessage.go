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
	ErrNone                             = New(100000, "成功")
	ErrDefault                          = New(1, "系统繁忙，稍后重试")
	ErrAesNewCipher                     = New(2, "")
	ErrBase64RawURLEncodingDecodeString = New(3, "")
	ErrGrpcDial                         = New(4, "")
	ErrPbNewServiceClientCall           = New(5, "")
	ErrCmdOutput                        = New(6, "")
	ErrHttpRequestParseForm             = New(7, "")
	ErrHttpRequestParseMultipartForm    = New(8, "")
	ErrJsonNewDecoderDecode             = New(9, "")
	ErrHttpRequestBodyClose             = New(10, "")
	ErrOsOpen                           = New(11, "")
	ErrCsvNewReaderRead                 = New(12, "")
	ErrCsvNewReaderReadAll              = New(13, "")
	ErrOsFileClose                      = New(14, "")
	ErrOsOpenFile                       = New(15, "")
	ErrCsvNewWriterWrite                = New(16, "")
	ErrOsStat                           = New(17, "")
	ErrOsMkdirAll                       = New(18, "")
	ErrOsRemoveAll                      = New(19, "")
	ErrOsRename                         = New(20, "")
	ErrOsIoCopy                         = New(21, "")
	ErrOsCreate                         = New(22, "")
	ErrOsFileSeek                       = New(22, "")
	ErrIoUtilReadAll                    = New(23, "")
	ErrOsReadFile                       = New(24, "")
	ErrOsFileWriteAt                    = New(25, "")
	ErrHashHashWrite                    = New(26, "")
	ErrJsonMarshal                      = New(27, "")
	ErrJsonUnmarshal                    = New(28, "")
	ErrGoMailDialerDialAndSend          = New(29, "")
	ErrSyscallMmap                      = New(30, "")
	ErrSyscallSyscall                   = New(31, "")
	ErrSyscallMunmap                    = New(32, "")
	ErrBcryptGenerateFromPassword       = New(33, "")
)
