package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_GetRtpInfo(t *testing.T) {
	resp, err := client.GetRtpInfo(context.Background(), &GetRtpInfoRequest{
		Secret:   "",
		Vhost:    zlm_def.Vhost_Default,
		App:      testApp,
		StreamId: testStream,
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp)
}
