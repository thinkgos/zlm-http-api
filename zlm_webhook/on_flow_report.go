package zlm_webhook

//* 流量统计事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_4%E3%80%81on-flow-report
//
// 播放器或推流器断开时并且耗用流量超过特定阈值时会触发此事件,
// 阈值通过配置文件general.flowThreshold配置; 此事件对回复不敏感.

type OnFlowReportRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Schema        string `json:"schema"`        // 播放或推流的协议, 可能是 rtsp、rtmp、http
	Vhost         string `json:"vhost"`         // 流虚拟主机
	App           string `json:"app"`           // 流应用名
	Stream        string `json:"stream"`        // 流id
	Params        string `json:"params"`        // 推流或播放url参数
	Protocol      string `json:"protocol"`      // 流协议
	Id            string `json:"id"`            // tcp链接唯一id
	Duration      int    `json:"duration"`      // tcp链接维持时间, 单位秒
	Player        bool   `json:"player"`        // true: 播放器, false: 推流器
	TotalBytes    int    `json:"totalBytes"`    // 耗费上下行流量总和, 单位字节
	Ip            string `json:"ip"`            // 客户端ip
	Port          int    `json:"port"`          // 客户端端口号
}
type OnFlowReportReply = DefaultReply
