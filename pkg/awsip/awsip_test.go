package awsip

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReading(t *testing.T) {
	ranges, err := ReadFile("ranges-1682907189.json")
	require.NoError(t, err)
	require.Greater(t, len(ranges.Prefixes), 0)
	require.Greater(t, len(ranges.Ipv6Prefixes), 0)
	require.NotEmpty(t, ranges.SyncToken)
	require.NotEmpty(t, ranges.CreateDate)
}
