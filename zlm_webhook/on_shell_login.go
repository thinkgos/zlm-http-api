package zlm_webhook

// * shell登录鉴权, telnet调试方式
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_11%E3%80%81on-shell-login
//
// 使用telnet 127.0.0.1 9000能进入MediaServer进程的shell界面。
type OnShellLoginRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Id            string `json:"id"`            // tcp链接唯一id
	Ip            string `json:"ip"`            // telnet终端ip
	Port          int    `json:"port"`          // telnet终端端口号
	UserName      string `json:"user_name"`     // telnet终端登录用户名
	Passwd        string `json:"passwd"`        // telnet终端登录用户密码
}

type OnShellLoginReply = DefaultReply
