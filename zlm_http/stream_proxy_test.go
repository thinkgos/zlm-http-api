package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

const (
	testStreamProxyVhost  = zlm_def.DefaultVhost
	testStreamProxyApp    = "live"
	testStreamProxyStream = "dahua"
	testStreamProxyUrl    = "rtsp://admin:admin123@10.110.18.75:554/cam/realmonitor?channel=1&subtype=0"
	testStreamProxyKey    = testStreamProxyVhost + "/" + testStreamProxyApp + "/" + testStreamProxyStream
)

func Test_ListStreamProxy(t *testing.T) {
	resp, err := client.ListStreamProxy(context.Background(), &ListStreamProxyRequest{})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_GetStreamProxyInfo(t *testing.T) {
	resp, err := client.GetStreamProxyInfo(context.Background(), &GetStreamProxyInfoRequest{
		Secret: "",
		Key:    testStreamProxyKey,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_AddStreamProxy(t *testing.T) {
	resp, err := client.AddStreamProxy(context.Background(), &AddStreamProxyRequest{
		Secret: "",
		Vhost:  testStreamProxyVhost,
		App:    testStreamProxyApp,
		Stream: testStreamProxyStream,
		Url:    testStreamProxyUrl,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_DelStreamProxy(t *testing.T) {
	resp, err := client.DelStreamProxy(context.Background(), &DelStreamProxyRequest{
		Secret: "",
		Key:    testStreamProxyKey,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}
