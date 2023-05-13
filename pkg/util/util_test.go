package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("valid uuid", func(t *testing.T) {
		require.True(t, IsValidUUID("fbd3036f-0f1c-4e98-b71c-d4cd61213f90"))
		require.True(t, IsValidUUID("123e4567-e89b-12d3-a456-426614174000"))
		require.True(t, IsValidUUID("ffffffff-ffff-ffff-ffff-ffffffffffff"))
		require.True(t, IsValidUUID("00000000-0000-0000-0000-000000000000"))
	})
	t.Run("invalid uuid", func(t *testing.T) {
		require.False(t, IsValidUUID(""))
		require.False(t, IsValidUUID("123"))
		require.False(t, IsValidUUID("f-bd3036f0f1c-4e98-b71c-d4cd61213f90"))
	})
}
