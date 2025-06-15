package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_GetMediaPlayerList(t *testing.T) {
	resp, err := client.GetMediaPlayerList(context.Background(), &GetMediaPlayerListRequest{
		Schema: testSchema,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.Data), 0)
	t.Logf("%#v\n", resp.Data)
}
