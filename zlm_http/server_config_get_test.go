package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetServerConfig(t *testing.T) {
	resp, err := client.GetServerConfig(context.Background(), &GetServerConfigRequest{})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.Data), 0)
	// t.Logf("%#v\n", resp.Data)
}
