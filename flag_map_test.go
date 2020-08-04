package feato_test

import (
	"testing"

	"github.com/imantung/feato"
	"github.com/stretchr/testify/require"
)

func TestFeatureStore_IsEnabled(t *testing.T) {
	store := feato.FlagMap{
		"name01": feato.Enabled,
		"name02": feato.Disabled,
		"name03": nil,
	}

	require.True(t, store.IsEnabled("name01"))
	require.False(t, store.IsEnabled("name02"))
	require.False(t, store.IsEnabled("name03"))
	require.False(t, store.IsEnabled("not-found"))
}
