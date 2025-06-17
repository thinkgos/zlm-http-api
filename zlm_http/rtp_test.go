package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_StartSendRtp(t *testing.T) {
	resp, err := client.StartSendRtp(context.Background(), &StartSendRtpRequest{
		Secret:  "",
		Vhost:   zlm_def.Vhost_Default,
		App:     testApp,
		Stream:  testStream,
		DstUrl:  testZLMediaKitIp,
		DstPort: 20008,
		IsUdp:   zlm_def.False,
		Ssrc:    1234567,
		SrcPort: nil,
		Type:    zlm_def.ToPtr(1),
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp)
}

func Test_StopSendRtp(t *testing.T) {
	resp, err := client.StopSendRtp(context.Background(), &StopSendRtpRequest{
		Secret: "",
		Vhost:  zlm_def.Vhost_Default,
		App:    testApp,
		Stream: testStream,
		Ssrc:   nil,
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
}
