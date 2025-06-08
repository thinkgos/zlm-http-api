package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetThreadsLoad(t *testing.T) {
	resp, err := client.GetThreadsLoad(context.Background(), &GetThreadsLoadRequest{})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.Data), 0)
	// t.Logf("%#v\n", resp.Data)
}
