package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

const (
	testFFmpegSourceVhost   = zlm_def.DefaultVhost
	testFFmpegSource_SrcUrl = "rtsp://admin:admin123@10.110.18.75:554/cam/realmonitor?channel=1&subtype=0"
	testFFmpegSource_DstUrl = "rtmp://" + testZLMediaKitIp + "/live/dahua"
	testFFMpegCmdKey        = "ffmpeg.cmd"
)

func Test_ListFFmpegSource(t *testing.T) {
	resp, err := client.ListFFmpegSource(context.Background(), &ListFFmpegSourceRequest{})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}
func Test_AddFFmpegSource(t *testing.T) {
	resp, err := client.AddFFmpegSource(context.Background(), &AddFFmpegSourceRequest{
		Secret:       "",
		SrcUrl:       testFFmpegSource_SrcUrl,
		DstUrl:       testFFmpegSource_DstUrl,
		FFmpegCmdKey: testFFMpegCmdKey,
		TimeoutMs:    10000,
		EnableHls:    0,
		EnableMp4:    0,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}
func Test_DelFFmpegSource(t *testing.T) {
	const testFFmpegSourceKey = ""

	resp, err := client.DelFFmpegSource(context.Background(), &DelFFmpegSourceRequest{
		Secret: "",
		Key:    testFFmpegSourceKey,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}

func Test_FFmpegSource(t *testing.T) {
	respAddFFmpegSource, err := client.AddFFmpegSource(context.Background(), &AddFFmpegSourceRequest{
		Secret:       "",
		SrcUrl:       testFFmpegSource_SrcUrl,
		DstUrl:       testFFmpegSource_DstUrl,
		FFmpegCmdKey: testFFMpegCmdKey,
		TimeoutMs:    10000,
		EnableHls:    0,
		EnableMp4:    0,
	})
	require.NoError(t, err)
	t.Logf("%#v\n", respAddFFmpegSource.Data)
	testFFmpegSourceKey := respAddFFmpegSource.Data.Key

	respListFFmpegSource, err := client.ListFFmpegSource(context.Background(), &ListFFmpegSourceRequest{})
	require.NoError(t, err)
	require.Equal(t, 1, len(respListFFmpegSource.Data))
	t.Logf("%#v\n", respListFFmpegSource.Data)

	respDelFFmpegSource, err := client.DelFFmpegSource(context.Background(), &DelFFmpegSourceRequest{
		Secret: "",
		Key:    testFFmpegSourceKey,
	})
	require.NoError(t, err)
	require.True(t, respDelFFmpegSource.Data.Flag)
	t.Logf("%#v\n", respDelFFmpegSource.Data)

	respListFFmpegSource1, err := client.ListFFmpegSource(context.Background(), &ListFFmpegSourceRequest{})
	require.NoError(t, err)
	require.Equal(t, 0, len(respListFFmpegSource1.Data))
}
