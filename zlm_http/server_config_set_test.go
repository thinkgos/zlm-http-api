package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

func Test_SetServerConfig(t *testing.T) {
	resp, err := client.SetServerConfig(context.Background(), &SetServerConfigRequest{
		Api_ApiDebug: zlm_def.ToPtr(1),
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp)
}
