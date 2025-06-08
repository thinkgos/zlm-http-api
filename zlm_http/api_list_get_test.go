package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetApiList(t *testing.T) {
	resp, err := client.GetApiList(context.Background(), &GetApiListRequest{})
	require.NoError(t, err)
	require.Greater(t, len(resp.Data), 0)
	// t.Logf("%#v\n", resp.Data)
}
