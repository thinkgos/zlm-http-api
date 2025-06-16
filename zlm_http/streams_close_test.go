package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_CloseStreams(t *testing.T) {
	resp, err := client.CloseStreams(context.Background(), &CloseStreamsRequest{
		Secret: "",
		Schema: testSchema,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Force:  zlm_def.ToPtr(1),
	})
	require.NoError(t, err)
	t.Logf("%#v\n", resp)
}
