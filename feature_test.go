package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/imantung/feato"
)

func TestFeature_New(t *testing.T) {
	feature := feato.NewFeature("some-id")
	require.Equal(t, feato.DisableIndexToggle, feature.DefaultIndexToggle)
	require.Equal(t, "some-id", feature.ID)
}

func TestFeature_NewUnique(t *testing.T) {
	feature1 := feato.NewUniqueFeature()
	feature2 := feato.NewUniqueFeature()
	require.NotEqual(t, feature1.ID, feature2.ID)
}

func TestFeature_SetDefaultIndexToggle(t *testing.T) {
	feature := feato.NewUniqueFeature().SetDefaultIndexToggle(99)
	require.Equal(t, feato.IndexToggle(99), feature.DefaultIndexToggle)
}
