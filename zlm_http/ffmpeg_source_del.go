package zlm_http

import "context"

//* 关闭 ffmpeg 拉流代, 流注册成功后，也可以使用close_streams接口替代
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_15%E3%80%81-index-api-delffmpegsource-%E6%B5%81%E6%B3%A8%E5%86%8C%E6%88%90%E5%8A%9F%E5%90%8E-%E4%B9%9F%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8close-streams%E6%8E%A5%E5%8F%A3%E6%9B%BF%E4%BB%A3

type DelFFmpegSourceRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Key    string `json:"key"`              // M, addFFmpegSource 接口返回的 key
}
type DelFFmpegSourceReply struct {
	BaseResult
	Data DelFFmpegSourceData `json:"data"`
}
type DelFFmpegSourceData struct {
	Flag bool `json:"flag"` // 成功与否
}

func (c *Client) DelFFmpegSource(ctx context.Context, req *DelFFmpegSourceRequest, opts ...CallOption) (*DelFFmpegSourceReply, error) {
	return genericPost[DelFFmpegSourceRequest, DelFFmpegSourceReply](c, "/index/api/delFFmpegSource", ctx, req, opts...)
}
