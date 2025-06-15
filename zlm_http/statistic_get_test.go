package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetStatistic(t *testing.T) {
	resp, err := client.GetStatistic(context.Background(), &GetStatisticRequest{})
	require.NoError(t, err)
	t.Logf("%#v\n", resp.Data)
}
