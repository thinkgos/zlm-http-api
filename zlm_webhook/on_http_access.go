package zlm_webhook

//* 访问http文件服务器上hls之外的文件时触发
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_5%E3%80%81on-http-access

type OnHttpAccessRequest struct {
	MediaServerId                 string `json:"mediaServerId"`                    // 服务器 id, 通过配置文件设置
	HeaderAccept                  string `json:"header.Accept"`                    // http客户端请求header
	HeaderAcceptEncoding          string `json:"header.Accept-Encoding"`           // http客户端请求header
	HeaderAcceptLanguage          string `json:"header.Accept-Language"`           // http客户端请求header
	HeaderCacheControl            string `json:"header.Cache-Control"`             // http客户端请求header
	HeaderConnection              string `json:"header.Connection"`                // http客户端请求header
	HeaderHost                    string `json:"header.Host"`                      // http客户端请求header
	HeaderUpgradeInsecureRequests string `json:"header.Upgrade-Insecure-Requests"` // http客户端请求header
	HeaderUserAgent               string `json:"header.User-Agent"`                // http客户端请求header
	Id                            string `json:"id"`                               // tcp链接唯一id
	Ip                            string `json:"ip"`                               // http客户端ip
	Port                          int    `json:"port"`                             // http客户端端口号
	IsDir                         bool   `json:"is_dir"`                           // http访问路径是文件还是目录
	Params                        string `json:"params"`                           // http url参数
	Path                          string `json:"path"`                             // 请求访问的文件或目录
}

type OnHttpAccessReply struct {
	Code          int    `json:"code"`          // 请固定返回0
	Err           string `json:"err"`           // 不允许访问的错误提示, 允许访问请置空
	Path          string `json:"path"`          // 该客户端能访问或被禁止的顶端目录, 如果为空字符串, 则表述为当前目录
	Second        int    `json:"second"`        // 本次授权结果的有效期, 单位秒
	MediaServerId string `json:"mediaServerId"` // 服务器 id, 通过配置文件设置

}
