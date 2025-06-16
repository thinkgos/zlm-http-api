package zlm_http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetAllSession(t *testing.T) {
	resp, err := client.GetAllSession(context.Background(), &GetAllSessionRequest{})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp.Data)
}

func Test_KickSession(t *testing.T) {
	resp, err := client.KickSession(context.Background(), &KickSessionRequest{
		Id: "123-108",
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp)
}

func Test_KickSessions(t *testing.T) {
	resp, err := client.KickSessions(context.Background(), &KickSessionsRequest{
		Secret:    "",
		LocalPort: 0,
		PeerIp:    "",
	})
	require.NoError(t, err)
	require.Equal(t, 0, resp.Code)
	t.Logf("%#v\n", resp)
}
