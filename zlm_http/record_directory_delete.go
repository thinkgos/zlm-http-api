package zlm_http

import "context"

//* 删除录像文件夹

type DeleteRecordDirectoryRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost  string `json:"vhost"`            // M, 流的虚拟主机, 例如__defaultVhost__
	App    string `json:"app"`              // M, 流的应用名, 例如 live
	Stream string `json:"stream"`           // M, 流id, 例如 test
	Period string `json:"period"`           // M, 流的录像日期，格式为2020-01-01,如果不是完整的日期，那么会删除失败
}
type DeleteRecordDirectoryReply struct {
	BaseResult
	Path string `json:"path"` // 录像文件夹路径
}

func (c *Client) DeleteRecordDirectory(ctx context.Context, req *DeleteRecordDirectoryRequest, opts ...CallOption) (*DeleteRecordDirectoryReply, error) {
	return GenericPost[DeleteRecordDirectoryRequest, DeleteRecordDirectoryReply](c, "/index/api/deleteRecordDirectory", ctx, req, opts...)
}
