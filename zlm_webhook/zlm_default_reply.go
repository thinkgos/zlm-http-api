package zlm_webhook

const (
	Code_Success     = 0
	Code_Failure     = 1
	Code_Success_Msg = "success"
)

type DefaultReply struct {
	Code int    `json:"code"` // 错误代码, 0: 成功
	Msg  string `json:"msg"`  // 错误提示
}

func NewDefaultSuccessReply() *DefaultReply {
	return NewDefaultReply(Code_Success, Code_Success_Msg)
}

func NewDefaultFailureReply(msg string) *DefaultReply {
	return NewDefaultReply(Code_Failure, msg)
}

func NewDefaultReply(code int, msg string) *DefaultReply {
	return &DefaultReply{
		Code: code,
		Msg:  msg,
	}
}
