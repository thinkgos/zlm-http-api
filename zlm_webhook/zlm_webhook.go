package zlm_webhook

//* doc: https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html

type DefaultReply struct {
	Code int    `json:"code"` // 错误代码, 0: 成功
	Msg  string `json:"msg"`  // 错误提示
}

func NewDefaultReply() DefaultReply {
	return DefaultReply{Code: 0, Msg: "success"}
}
