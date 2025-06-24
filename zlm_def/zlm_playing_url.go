package zlm_def

import (
	"net/url"
	"strconv"
	"strings"
)

// 播放url
type PlayingUrl struct {
	App         string
	Stream      string
	Host        string
	HttpPort    int
	HttpSslPort int
	RtmpPort    int
	RtmpSslPort int
	RtspPort    int
	RtspSslPort int
	Param       url.Values
}

func (p *PlayingUrl) build(schema, realPort, realStream string) string {
	b := strings.Builder{}
	b.WriteString(schema)
	b.WriteString("://")
	b.WriteString(p.Host)
	if realPort != "" {
		b.WriteString(":")
		b.WriteString(realPort)
	}
	b.WriteString("/")
	b.WriteString(p.App)
	b.WriteString("/")
	b.WriteString(realStream)
	if params := p.Param.Encode(); params != "" {
		b.WriteString("?")
		b.WriteString(params)
	}
	return b.String()
}

func (p *PlayingUrl) Rtsp() string {
	return p.build("rtsp", realPort(p.RtspPort, 554), p.Stream)
}

func (p *PlayingUrl) Rtsps() string {
	return p.build("rtsps", realPort(p.RtspSslPort, 0), p.Stream)
}

func (p *PlayingUrl) Rtmp() string {
	return p.build("rtmp", realPort(p.RtmpPort, 1935), p.Stream)
}

func (p *PlayingUrl) Rtmps() string {
	return p.build("rtmps", realPort(p.RtmpSslPort, 0), p.Stream)
}

func (p *PlayingUrl) HttpFlv() string {
	return p.build("http", realPort(p.HttpPort, 80), p.Stream+".live.flv")
}

func (p *PlayingUrl) HttpsFlv() string {
	return p.build("https", realPort(p.HttpSslPort, 443), p.Stream+".live.flv")
}

// ws-flv
func (p *PlayingUrl) WsFlv() string {
	return p.build("ws", realPort(p.HttpPort, 80), p.Stream+".live.flv")
}

// wss-flv
func (p *PlayingUrl) WssFlv() string {
	return p.build("wss", realPort(p.HttpSslPort, 443), p.Stream+".live.flv")
}

// http-HLS(mpegts)
func (p *PlayingUrl) HttpHls() string {
	return p.build("http", realPort(p.HttpPort, 80), p.Stream+"/hls.m3u8")
}

// http-HLS(mpegts)
func (p *PlayingUrl) HttpsHls() string {
	return p.build("https", realPort(p.HttpSslPort, 443), p.Stream+"/hls.m3u8")
}

// http-HLS(fmp4)
func (p *PlayingUrl) HttpHlsFmp4() string {
	return p.build("http", realPort(p.HttpPort, 80), p.Stream+"/hls.fmp4.m3u8")
}

// http-HLS(fmp4)
func (p *PlayingUrl) HttpsHlsFmp4() string {
	return p.build("https", realPort(p.HttpSslPort, 443), p.Stream+"/hls.fmp4.m3u8")
}

// http-ts
func (p *PlayingUrl) HttpTs() string {
	return p.build("http", realPort(p.HttpPort, 80), p.Stream+".live.ts")
}

// https-ts
func (p *PlayingUrl) HttpsTs() string {
	return p.build("https", realPort(p.HttpSslPort, 443), p.Stream+".live.ts")
}

// ws-ts
func (p *PlayingUrl) WsTs() string {
	return p.build("ws", realPort(p.HttpPort, 80), p.Stream+".live.ts")
}

// wss-ts
func (p *PlayingUrl) WssTs() string {
	return p.build("wss", realPort(p.HttpSslPort, 443), p.Stream+".live.ts")
}

// http-fmp4
func (p *PlayingUrl) HttpFmp4() string {
	return p.build("http", realPort(p.HttpPort, 80), p.Stream+".live.mp4")
}

// https-fmp4
func (p *PlayingUrl) HttpsFmp4() string {
	return p.build("https", realPort(p.HttpSslPort, 443), p.Stream+".live.mp4")
}

// ws-fmp4
func (p *PlayingUrl) WsFmp4() string {
	return p.build("ws", realPort(p.HttpPort, 80), p.Stream+".live.mp4")
}

// wss-fmp4
func (p *PlayingUrl) WssFmp4() string {
	return p.build("wss", realPort(p.HttpSslPort, 443), p.Stream+".live.mp4")
}

func realPort(port, defaultPort int) string {
	if port > 0 && (defaultPort <= 0 || port != defaultPort) {
		return strconv.Itoa(port)
	}
	return ""
}
