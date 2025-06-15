package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_IsRecording(t *testing.T) {
	resp, err := client.IsRecording(context.Background(), &IsRecordingRequest{
		Secret: testSecret,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Type:   1,
	})
	require.NoError(t, err)
	// require.True(t, resp.Status)
	require.False(t, resp.Status)
	t.Logf("%#v\n", resp)
}
