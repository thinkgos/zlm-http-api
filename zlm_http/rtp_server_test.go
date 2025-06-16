package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RtpServer(t *testing.T) {
	//* 查询rtp流, 0个
	respListRtpServerZero, err := client.ListRtpServer(context.Background(), &ListRtpServerRequest{})
	require.NoError(t, err)
	require.Equal(t, 0, len(respListRtpServerZero.Data))

	//* 创建rtp流
	respOpenRtpServer, err := client.OpenRtpServer(context.Background(), &OpenRtpServerRequest{
		// Vhost:     zlm_def.DefaultVhost,
		// App:       testApp,
		StreamId:  testStream,
		TcpMode:   1,
		Port:      0,
		OnlyTrack: 0,
		Ssrc:      0,
		LocalIp:   "",
		ReUsePort: false,
	})
	require.NoError(t, err)
	require.Equal(t, 0, respOpenRtpServer.Code)
	require.Greater(t, respOpenRtpServer.Port, 0)

	//* 查询rtp流, 1个
	respListRtpServerOne, err := client.ListRtpServer(context.Background(), &ListRtpServerRequest{})
	require.NoError(t, err)
	require.Equal(t, 1, len(respListRtpServerOne.Data))
	t.Logf("%#v\n", respListRtpServerOne.Data)

	//* 关闭rtp流
	respCloseRtpServer, err := client.CloseRtpServer(context.Background(), &CloseRtpServerRequest{
		// Vhost:    zlm_def.DefaultVhost,
		// App:      testApp,
		StreamId: testStream,
	})
	require.NoError(t, err)
	require.Equal(t, 0, respCloseRtpServer.Code)
	require.Equal(t, 1, respCloseRtpServer.Hit)

	//* 查询rtp流, 0个
	respListRtpServerZero1, err := client.ListRtpServer(context.Background(), &ListRtpServerRequest{})
	require.NoError(t, err)
	require.Equal(t, 0, len(respListRtpServerZero1.Data))
}
