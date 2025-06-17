package zlm_http

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Download(t *testing.T) {
	const testFilePath = "/opt/media/bin/www/record/live/chibo01/2025-06-16/2025-06-16-16-53-10-0.mp4"

	defer func() {
		_ = os.Remove(filepath.Base(testFilePath))
	}()

	err := client.DownloadFile(context.Background(), &DownloadFileRequest{
		FilePath: testFilePath,
		SaveFile: "",
	})
	require.NoError(t, err)
}
