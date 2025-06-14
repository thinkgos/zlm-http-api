package zlm_webhook

import "github.com/thinkgos/zlm-http-api/zlm_def"

//* 服务器启动事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_15%E3%80%81on-server-started
//
// 可以用于监听服务器崩溃重启; 此事件对回复不敏感.

type OnServerStartedRequest struct {
	HookIndex int `json:"hook_index"`
	zlm_def.ServerConfig
}
type OnServerStartedReply = DefaultReply
