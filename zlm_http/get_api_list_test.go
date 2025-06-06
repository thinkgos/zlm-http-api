package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetApiList(t *testing.T) {
	resp, err := client.GetApiList(context.Background(), &GetAPIListParams{})
	require.NoError(t, err)
	t.Log(resp)
}
