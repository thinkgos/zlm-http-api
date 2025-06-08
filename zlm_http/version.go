package zlm_http

import "context"

//* 获取版本信息, 如分支, commit id, 编译时间
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_32%E3%80%81-index-api-version-%E8%8E%B7%E5%8F%96%E7%89%88%E6%9C%AC%E4%BF%A1%E6%81%AF

type VersionRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type VersionReply struct {
	BaseResult
	Data *VersionData `json:"data"`
}
type VersionData struct {
	BranchName string `json:"branchName"` // 分支
	BuildTime  string `json:"buildTime"`  // 编译时间
	CommitHash string `json:"commitHash"` // commit id
}

func (c *Client) Version(ctx context.Context, req *VersionRequest, opts ...CallOption) (*VersionReply, error) {
	var resp VersionReply

	err := c.Get(ctx, "/index/api/version", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
