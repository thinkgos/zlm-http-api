package zlm_http

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thinkgos/zlm-http-api/zlm_def"
)

const testRecordType = 0 // 0: hls, 1: mp4

func Test_Record(t *testing.T) {
	respStartRecord, err := client.StartRecord(context.Background(), &StartRecordRequest{
		Secret:         testSecret,
		Vhost:          zlm_def.DefaultVhost,
		App:            testApp,
		Stream:         testStream,
		Type:           testRecordType,
		CustomizedPath: "",
		MaxSecond:      5,
	})
	require.NoError(t, err)
	require.True(t, respStartRecord.Result)
	t.Logf("%#v\n", respStartRecord)

	respIsRecordingTrue, err := client.IsRecording(context.Background(), &IsRecordingRequest{
		Secret: testSecret,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Type:   testRecordType,
	})
	require.NoError(t, err)
	require.True(t, respIsRecordingTrue.Status)
	t.Logf("%#v\n", respIsRecordingTrue)

	time.Sleep(time.Second)

	respStopRecord, err := client.StopRecord(context.Background(), &StopRecordRequest{
		Secret: testSecret,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Type:   testRecordType,
	})
	require.NoError(t, err)
	require.True(t, respStopRecord.Result)
	t.Logf("%#v\n", respStopRecord)

	respIsRecordingFalse, err := client.IsRecording(context.Background(), &IsRecordingRequest{
		Secret: testSecret,
		Vhost:  zlm_def.DefaultVhost,
		App:    testApp,
		Stream: testStream,
		Type:   testRecordType,
	})
	require.NoError(t, err)
	require.False(t, respIsRecordingFalse.Status)
	t.Logf("%#v\n", respIsRecordingFalse)

	if testRecordType == 1 {
		respRecordFile, err := client.GetMp4RecordFile(context.Background(), &GetMp4RecordFileRequest{
			Secret: testSecret,
			Vhost:  zlm_def.DefaultVhost,
			App:    testApp,
			Stream: testStream,
			Period: "2025-06",
		})
		require.NoError(t, err)
		require.Equal(t, 0, respRecordFile.Code)
		t.Logf("%#v\n", respRecordFile.Data)
	}
}
