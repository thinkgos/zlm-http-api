package zlm_http

import "context"

//* 关闭 GB28181 RTP 接收端口
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_25%E3%80%81-index-api-closertpserver

type CloseRtpServerRequest struct {
	StreamId string `json:"stream_id"` // M, 该端口绑定的流id, 该端口只能创建这一个流(而不是根据ssrc创建多个)
}
type CloseRtpServerReply struct {
	BaseResult
	Hit int `json:"hit"` // 是否找到记录并关闭
}

func (c *ZlmClient) CloseRtpServer(ctx context.Context, req *CloseRtpServerRequest) (*CloseRtpServerReply, error) {
	var resp CloseRtpServerReply

	err := c.Post(ctx, "/index/api/closeRtpServer", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
