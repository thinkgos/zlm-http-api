package zlm_def

import (
	"net/url"
	"testing"
)

func Test_PlayingUrl(t *testing.T) {
	p := PlayingUrl{
		App:    "live",
		Stream: "chibo01",
		// Host:        "127.0.0.1",
		Host:        "nova.thinkgos.cn",
		HttpPort:    80,
		HttpSslPort: 443,
		RtmpPort:    1935,
		RtmpSslPort: 0,
		RtspPort:    554,
		RtspSslPort: 555,
		Param:       url.Values{"token": []string{"abc123"}},
	}

	t.Log(p.Rtsp())
	t.Log(p.Rtsps())
	t.Log(p.Rtmp())
	t.Log(p.Rtmps())

	t.Log("--------------------------------")
	t.Log(p.HttpFlv())
	t.Log(p.HttpsFlv())
	t.Log(p.WsFlv())
	t.Log(p.WssFlv())

	t.Log("--------------------------------")
	t.Log(p.HttpHls())
	t.Log(p.HttpsHls())
	t.Log(p.HttpHlsFmp4())
	t.Log(p.HttpsHlsFmp4())

	t.Log("--------------------------------")
	t.Log(p.HttpTs())
	t.Log(p.HttpsTs())
	t.Log(p.WsTs())
	t.Log(p.WssTs())

	t.Log("--------------------------------")
	t.Log(p.HttpFmp4())
	t.Log(p.HttpsFmp4())
	t.Log(p.WsFmp4())
	t.Log(p.WssFmp4())
}
