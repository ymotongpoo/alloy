package database_observability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/grafana/alloy/internal/runtime/logging"
	"github.com/grafana/loki/pkg/push"
)

func TestBuildLokiEntry(t *testing.T) {
	entry := BuildLokiEntry(logging.LevelDebug, "test-operation", "This is a test log line")

	require.Len(t, entry.Labels, 1)
	require.Equal(t, "test-operation", string(entry.Labels["op"]))
	require.Equal(t, `level="debug" This is a test log line`, entry.Line)
}

func TestBuildLokiEntryWithTimestamp(t *testing.T) {
	entry := BuildLokiEntryWithTimestamp(logging.LevelInfo, "test-operation", "This is a test log line", 5)

	require.Equal(t, int64(5), entry.Entry.Timestamp.UnixNano())
	require.Equal(t, time.Unix(0, 5), entry.Entry.Timestamp)
}

func TestBuildLokiEntryWithStructuredMetadataAndTimestamp(t *testing.T) {
	metadata := push.LabelsAdapter{{Name: "wait_event_type", Value: "IO Wait"}}
	entry := BuildLokiEntryWithStructuredMetadataAndTimestamp(logging.LevelInfo, "test-operation", "This is a test log line", metadata, 5)

	require.Equal(t, "test-operation", string(entry.Labels["op"]))
	require.Equal(t, `level="info" This is a test log line`, entry.Line)
	require.Equal(t, time.Unix(0, 5), entry.Entry.Timestamp)
	require.Len(t, entry.StructuredMetadata, 1)
	require.Equal(t, "wait_event_type", entry.StructuredMetadata[0].Name)
	require.Equal(t, "IO Wait", entry.StructuredMetadata[0].Value)
}
