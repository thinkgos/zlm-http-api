package zlm_http

import (
	"fmt"
)

// code定义
const (
	Code_Success     = 0    // 执行成功
	Code_OtherFailed = -1   // 业务代码执行失败
	Code_AuthFailed  = -100 // 鉴权失败
	Code_SqlFailed   = -200 // sql执行失败
	Code_InvalidArgs = -300 // 参数不合法
	Code_Exception   = -400 // 代码抛异常
)

type BaseResult struct {
	Code   int    `json:"code"`   // 执行结果, 0: 成功, 其它为失败
	Msg    string `json:"msg"`    // 失败消息
	Result int    `json:"result"` // 业务代码执行失败具体原因
}

func (b BaseResult) inspectError() error {
	if b.Code == Code_Success {
		return nil
	}
	return &ErrorBusiness{
		Code:   b.Code,
		Msg:    b.Msg,
		Result: b.Result,
	}
}

type ErrorBusiness struct {
	Code   int    `json:"code"`   // 执行结果, 0: 成功, 其它为失败
	Msg    string `json:"msg"`    // 失败消息
	Result int    `json:"result"` // 业务代码执行失败具体原因
}

func (e *ErrorBusiness) Error() string {
	if e.Code == Code_Success {
		return "<nil>"
	}
	if e.Result == 0 {
		return fmt.Sprintf("zlm_http: code %d, msg %s", e.Code, e.Msg)
	}
	return fmt.Sprintf("zlm_http: code %d, msg %s, result %d", e.Code, e.Msg, e.Code)
}
