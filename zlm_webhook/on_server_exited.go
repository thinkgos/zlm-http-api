package zlm_webhook

//* 服务器退出事件

type OnServerExitedRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"`
}
type OnServerExitedReply = DefaultReply
