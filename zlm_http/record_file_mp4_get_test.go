package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_GetMp4RecordFile(t *testing.T) {
	resp, err := client.GetMp4RecordFile(context.Background(), &GetMp4RecordFileRequest{
		Secret: testSecret,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Period: "2025-06",
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp.Data)
}
