package zlm_http

const (
	CodeSuccess     = 0    // 执行成功
	CodeOtherFailed = -1   // 业务代码执行失败
	CodeAuthFailed  = -100 // 鉴权失败
	CodeSqlFailed   = -200 // sql执行失败
	CodeInvalidArgs = -300 // 参数不合法
	CodeException   = -400 // 代码抛异常
)

type BaseResult struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result int    `json:"result"`
}
