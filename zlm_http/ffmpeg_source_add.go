package zlm_http

import "context"

//* 通过 fork FFmpeg 进程的方式拉流代理，支持任意协议
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_14%E3%80%81-index-api-AddFFmpegSource

type AddFFmpegSourceRequest struct {
	Secret       string `json:"secret,omitempty"`  // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	SrcUrl       string `json:"src_url"`           // FFmpeg 拉流地址,支持任意协议或格式(只要 FFmpeg 支持即可)
	DstUrl       string `json:"dst_url"`           // FFmpeg rtmp 推流地址，一般都是推给自己，例如 rtmp://127.0.0.1/live/stream_form_ffmpeg
	TimeoutMs    int    `json:"timeout_ms,string"` // FFmpeg 推流成功超时时间
	EnableHls    int    `json:"enable_hls,string"` // 是否开启 hls 录制
	EnableMp4    int    `json:"enable_mp4,string"` // 是否开启 mp4 录制
	FfmpegCmdKey string `json:"ffmpeg_cmd_key"`    // 配置文件中 FFmpeg 命令参数模板 key(非内容)，置空则采用默认模板:ffmpeg.cmd
}
type AddFFmpegSourceReply struct {
	BaseResult
	Data AddFFmpegSourceData `json:"data"`
}
type AddFFmpegSourceData struct {
	Key string `json:"key"` // 唯一key
}

func (c *Client) AddFFmpegSource(ctx context.Context, req *AddFFmpegSourceRequest, opts ...CallOption) (*AddFFmpegSourceReply, error) {
	return GenericPost[AddFFmpegSourceRequest, AddFFmpegSourceReply](c, "/index/api/addFFmpegSource", ctx, req, opts...)
}
