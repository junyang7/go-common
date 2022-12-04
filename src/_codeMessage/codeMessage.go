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
	ErrParameter                        = New(2, "参数错误")
	ErrAesNewCipher                     = New(101, "")
	ErrBase64RawURLEncodingDecodeString = New(102, "")
	ErrGrpcDial                         = New(103, "")
	ErrPbNewServiceClientCall           = New(104, "")
	ErrCmdOutput                        = New(105, "")
	ErrHttpRequestParseForm             = New(106, "")
	ErrHttpRequestParseMultipartForm    = New(107, "")
	ErrJsonNewDecoderDecode             = New(108, "")
	ErrHttpRequestBodyClose             = New(109, "")
	ErrOsOpen                           = New(110, "")
	ErrCsvNewReaderRead                 = New(111, "")
	ErrCsvNewReaderReadAll              = New(112, "")
	ErrOsFileClose                      = New(113, "")
	ErrOsOpenFile                       = New(114, "")
	ErrCsvNewWriterWrite                = New(115, "")
	ErrOsStat                           = New(116, "")
	ErrOsMkdirAll                       = New(117, "")
	ErrOsRemoveAll                      = New(118, "")
	ErrOsRename                         = New(119, "")
	ErrOsIoCopy                         = New(120, "")
	ErrOsCreate                         = New(121, "")
	ErrOsFileSeek                       = New(122, "")
	ErrIoUtilReadAll                    = New(123, "")
	ErrOsReadFile                       = New(124, "")
	ErrOsFileWriteAt                    = New(125, "")
	ErrHashHashWrite                    = New(126, "")
	ErrJsonMarshal                      = New(127, "")
	ErrJsonUnmarshal                    = New(128, "")
	ErrGoMailDialerDialAndSend          = New(129, "")
	ErrSyscallMmap                      = New(130, "")
	ErrSyscallSyscall                   = New(131, "")
	ErrSyscallMunmap                    = New(132, "")
	ErrBcryptGenerateFromPassword       = New(133, "")
)
