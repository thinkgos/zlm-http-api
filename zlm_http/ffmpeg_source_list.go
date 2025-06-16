package zlm_http

import "context"

type ListFFmpegSourceRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type ListFFmpegSourceReply struct {
	BaseResult
	Data []FFmpegSourceEntry `json:"data"`
}
type FFmpegSourceEntry struct {
	Key          string `json:"key"`
	SrcURL       string `json:"src_url"`
	DstURL       string `json:"dst_url"`
	FFmpegCmdKey string `json:"ffmpeg_cmd_key"`
	Cmd          string `json:"cmd"`
}

func (c *Client) ListFFmpegSource(ctx context.Context, req *ListFFmpegSourceRequest, opts ...CallOption) (*ListFFmpegSourceReply, error) {
	return GenericPost[ListFFmpegSourceRequest, ListFFmpegSourceReply](c, "/index/api/listFFmpegSource", ctx, req, opts...)
}
