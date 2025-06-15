package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RestartServer(t *testing.T) {
	resp, err := client.RestartServer(context.Background(), &RestartServerRequest{})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp.Msg)
}
