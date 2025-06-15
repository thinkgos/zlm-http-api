package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_StopRecord(t *testing.T) {
	resp, err := client.StopRecord(context.Background(), &StopRecordRequest{
		Secret: testSecret,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Type:   1,
	})
	require.NoError(t, err)
	require.True(t, resp.Result)
	t.Logf("%#v\n", resp)
}
