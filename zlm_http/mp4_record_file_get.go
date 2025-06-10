package zlm_http

import "context"

//* 搜索文件系统，获取流对应的录像文件列表或日期文件夹列表
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_19%E3%80%81-index-api-getmp4recordfile

type GetMp4RecordFileRequest struct {
	Secret         string `json:"secret,omitempty"`          // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost          string `json:"vhost"`                     // M, 流的虚拟主机, 例如__defaultVhost__
	App            string `json:"app"`                       // M, 流的应用名, 例如 live
	Stream         string `json:"stream"`                    // M, 流id, 例如 test
	Period         string `json:"period"`                    // M, 流的录像日期, 格式为2020-02-01/2020-02, 如果不是完整的日期，那么是搜索录像文件夹列表，否则搜索对应日期下的 mp4 文件列表
	CustomizedPath string `json:"customized_path,omitempty"` // O, 自定义搜索路径，与 startRecord 方法中的 customized_path 一样，默认为配置文件的路径
}
type GetMp4RecordFileReply struct {
	BaseResult
	Data *GetMp4RecordFileData `json:"data"`
}
type GetMp4RecordFileData struct {
	Paths    []string `json:"paths"`
	RootPath string   `json:"rootPath"`
}

func (c *Client) GetMp4RecordFile(ctx context.Context, req *GetMp4RecordFileRequest, opts ...CallOption) (*GetMp4RecordFileReply, error) {
	return genericGet[GetMp4RecordFileRequest, GetMp4RecordFileReply](c, "/index/api/getMp4RecordFile", ctx, req, opts...)
}
