package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_StartRecord(t *testing.T) {
	resp, err := client.StartRecord(context.Background(), &StartRecordRequest{
		Secret:         testSecret,
		Vhost:          zlm_def.DefaultVhost,
		App:            testApp,
		Stream:         testStream,
		Type:           1,
		CustomizedPath: "",
		MaxSecond:      10,
	})
	require.NoError(t, err)
	require.True(t, resp.Result)
	t.Logf("%#v\n", resp)
}
