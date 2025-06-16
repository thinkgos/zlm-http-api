package zlm_http

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetSnap(t *testing.T) {
	const testFilename = "./test.jpg"

	defer os.Remove(testFilename)
	err := client.GetSnap(context.Background(), &GetSnapRequest{
		Url:        "12312",
		TimeoutSec: 10,
		ExpireSec:  10,
		Filename:   testFilename,
	})
	require.NoError(t, err)
}
