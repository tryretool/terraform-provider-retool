package acctest

import (
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

// This file borrows heavily from Auth0's Terraform provider: https://github.com/auth0/terraform-provider-auth0/blob/main/internal/acctest/http_recorder.go

const (
	recordingsDIR = "test/data/recordings/"

	// RecordingsHost is used for testing with our recorded http interactions.
	// The idea is that no matter which domain you're using to record the requests, the playback will always use this domain.
	RecordingsHost   = "recorded.retool.dev"
	RecordingsScheme = "https"
)

// NewHTTPRecorder creates a new instance of our http recorder used in tests.
func newHTTPRecorder(t *testing.T) *recorder.Recorder {
	t.Helper()

	recorderTransport, err := recorder.NewWithOptions(
		&recorder.Options{
			CassetteName:       cassetteName(t.Name()),
			Mode:               recorder.ModeRecordOnce,
			SkipRequestLatency: true,
		},
	)
	require.NoError(t, err)

	removeSensitiveDataFromRecordings(t, recorderTransport)

	t.Cleanup(func() {
		err := recorderTransport.Stop()
		require.NoError(t, err)
	})

	return recorderTransport
}

func cassetteName(testName string) string {
	_, file, _, _ := runtime.Caller(0)
	rootDir := path.Join(path.Dir(file), "../..")
	return path.Join(rootDir, recordingsDIR, testName)
}

func removeSensitiveDataFromRecordings(t *testing.T, recorderTransport *recorder.Recorder) {
	recorderTransport.AddHook(
		func(i *cassette.Interaction) error {
			skip429Response(i)
			redactHeaders(i)

			host := os.Getenv("RETOOL_HOST")
			require.NotEmpty(t, host, "removeSensitiveDataFromRecordings(): RETOOL_HOST is empty")
			scheme := os.Getenv("RETOOL_SCHEME")
			if scheme == "" {
				scheme = "https"
			}

			redactHostAndScheme(i, host, scheme)

			return nil
		},
		recorder.BeforeSaveHook,
	)
}

func skip429Response(i *cassette.Interaction) {
	if i.Response.Code == http.StatusTooManyRequests {
		i.DiscardOnSave = true
	}
}

func redactHeaders(i *cassette.Interaction) {
	allowedHeaders := map[string]bool{
		"Content-Type": true,
		"User-Agent":   true,
	}

	for header := range i.Request.Headers {
		if _, ok := allowedHeaders[header]; !ok {
			delete(i.Request.Headers, header)
		}
	}
	for header := range i.Response.Headers {
		if _, ok := allowedHeaders[header]; !ok {
			delete(i.Response.Headers, header)
		}
	}
}

func redactHostAndScheme(i *cassette.Interaction, host string, scheme string) {
	i.Request.Host = strings.ReplaceAll(i.Request.Host, host, RecordingsHost)
	i.Request.URL = strings.ReplaceAll(i.Request.URL, scheme+"://"+host, RecordingsScheme+"://"+RecordingsHost)
}
