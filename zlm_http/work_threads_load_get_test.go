package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetWorkThreadsLoad(t *testing.T) {
	resp, err := client.GetWorkThreadsLoad(context.Background(), &GetWorkThreadsLoadRequest{})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.Data), 0)
	// t.Logf("%#v\n", resp.Data)
}
