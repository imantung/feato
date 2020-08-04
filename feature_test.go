package feato_test

import (
	"testing"

	"github.com/imantung/feato"
	"github.com/stretchr/testify/require"
)

func TestFeature_New(t *testing.T) {
	feature := feato.NewFeature("some-id")
	require.Equal(t, feato.Disabled, feature.DefaultIndex)
	require.Equal(t, "some-id", feature.ID)
}

func TestFeature_SetDefaultIndexToggle(t *testing.T) {
	feature := feato.NewFeature("key1").SetDefaultIndex(99)
	require.Equal(t, feato.Index(99), feature.DefaultIndex)
}
