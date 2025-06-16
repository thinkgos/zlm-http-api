package zlm_http

import (
	"context"
	"net/http"
	"path/filepath"
)

//* 下载文件

type DownloadFileRequest struct {
	FilePath string `json:"file_path"` // M, 文件绝对路径，根据文件名生成Content-Type；该接口将触发on_http_access hook
	SaveFile string `json:"save_file"` // O, 浏览器下载文件后保存文件名；可选参数
}

func (c *Client) DownloadFile(ctx context.Context, req *DownloadFileRequest, opts ...CallOption) error {
	saveFilename := req.SaveFile
	if saveFilename == "" {
		saveFilename = filepath.Base(req.FilePath)
	}
	return c.Download(ctx, http.MethodGet, "/index/api/downloadFile", req, saveFilename, opts...)
}
