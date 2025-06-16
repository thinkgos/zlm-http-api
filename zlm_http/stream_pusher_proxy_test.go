package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

const (
	testStreamPusherProxySchema    = zlm_def.Schema_RTMP
	testStreamPusherProxyVhost     = zlm_def.DefaultVhost
	testStreamPusherProxyApp       = "live"
	testStreamPusherProxySrcStream = testStream
	testStreamPusherProxyDstStream = "chibo02"
	testStreamPusherProxyDstUrl    = testStreamPusherProxySchema + "://" + testZLMediaKitIp + "/" + testStreamPusherProxyApp + "/" + testStreamPusherProxyDstStream
	testStreamPusherProxyKey       = testStreamPusherProxySchema + "/" +
		testStreamPusherProxyVhost + "/" +
		testStreamPusherProxyApp + "/" +
		testStreamPusherProxySrcStream + "/" +
		"b0435dc2b84b886cba31ab5de9dbcfaa"
)

func Test_ListStreamPusherProxy(t *testing.T) {
	resp, err := client.ListStreamPusherProxy(context.Background(), &ListStreamPusherProxyRequest{})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_GetStreamPusherProxyInfo(t *testing.T) {
	resp, err := client.GetStreamPusherProxyInfo(context.Background(), &GetStreamPusherProxyInfoRequest{
		Secret: "",
		Key:    testStreamPusherProxyKey,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_AddStreamPusherProxy(t *testing.T) {
	resp, err := client.AddStreamPusherProxy(context.Background(), &AddStreamPusherProxyRequest{
		Secret: "",
		Vhost:  testStreamPusherProxyVhost,
		App:    testStreamPusherProxyApp,
		Stream: testStreamPusherProxySrcStream,
		DstUrl: testStreamPusherProxyDstUrl,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_DelStreamPusherProxy(t *testing.T) {
	resp, err := client.DelStreamPusherProxy(context.Background(), &DelStreamPusherProxyRequest{
		Secret: "",
		Key:    testStreamPusherProxyKey,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}
