package zlm_webhook

//* 录制mp4完成后通知事件; 此事件对回复不敏感.
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_8%E3%80%81on-record-mp4

type OnRecordMp4Request struct {
	HookIndex     int     `json:"hook_index"`
	MediaServerId string  `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	App           string  `json:"app"`           // 录制的流应用名
	Stream        string  `json:"stream"`        // 录制的流id
	Vhost         string  `json:"vhost"`         // 流虚拟主机
	Url           string  `json:"url"`           // http/rtsp/rtmp 点播相对url路径
	Params        string  `json:"params"`        // url参数
	Folder        string  `json:"folder"`        // 文件所在目录路径, 例: /opt/media/bin/www/live/chibo01
	FilePath      string  `json:"file_path"`     // 文件绝对路径, 例: /opt/media/bin/www/live/chibo01/2025-06-15/15/51-28_18.ts
	FileName      string  `json:"file_name"`     // 文件名, 例: 2025-06-15/15/51-28_18.ts
	FileSize      int     `json:"file_size"`     // 文件大小
	StartTime     int     `json:"start_time"`    // 开始录制时间戳, 单位秒
	TimeLen       float64 `json:"time_len"`      // 录制时长，单位秒
}
type OnRecordMp4Reply = DefaultReply
