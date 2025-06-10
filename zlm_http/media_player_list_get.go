package zlm_http

import "context"

//* 获取某个流观看者列表
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_33%E3%80%81-index-api-getmediaplayerlist

type GetMediaPlayerListRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Schema string `json:"schema"`           // M, 协议, 例如rtsp或 rtmp
	Vhost  string `json:"vhost"`            // M, 虚拟主机, 例如__defaultVhost__
	App    string `json:"app"`              // M, 应用名, 例如 live
	Stream string `json:"stream"`           // M, 流id, 例如 test
}
type GetMediaPlayerListReply struct {
	BaseResult
	Data []MediaPlayerEntry `json:"data"`
}

type MediaPlayerEntry struct {
	Identifier string `json:"identifier"`
	LocalIp    string `json:"local_ip"`
	LocalPort  int    `json:"local_port"`
	PeerIp     string `json:"peer_ip"`
	PeerPort   int    `json:"peer_port"`
	Typeid     string `json:"typeid"`
}

func (c *Client) GetMediaPlayerList(ctx context.Context, req *GetMediaPlayerListRequest, opts ...CallOption) (*GetMediaPlayerListReply, error) {
	var resp GetMediaPlayerListReply

	err := c.Get(ctx, "/index/api/getMediaPlayerList", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
