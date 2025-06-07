package zlm_http

import "context"

//* 获取 openRtpServer 接口创建的所有 RTP 服务器
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_26%E3%80%81-index-api-listrtpserver

type ListRtpServerRequest struct{}
type ListRtpServerReply struct {
	BaseResult
	Data []*RtpServerEntry `json:"data"` 
}
type RtpServerEntry struct {
	Port     int    `json:"port"`      // 绑定的端口号
	StreamId string `json:"stream_id"` //绑定的流id
}

func (c *ZlmClient) ListRtpServer(ctx context.Context, req *ListRtpServerRequest) (*ListRtpServerReply, error) {
	var resp ListRtpServerReply

	err := c.Get(ctx, "/index/api/listRtpServer", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
