package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetMediaList(t *testing.T) {
	resp, err := client.GetMediaList(context.Background(), &GetMediaListRequest{})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.Data), 0)
	// t.Logf("%#v\n", resp.Data)
}
