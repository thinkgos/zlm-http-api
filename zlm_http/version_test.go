package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Version(t *testing.T) {
	resp, err := client.Version(context.Background(), &VersionRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp.Data)
	t.Logf("%#v\n", resp.Data)
}
