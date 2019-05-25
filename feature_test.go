package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/feato"
)

func TestFeature_NewWithID(t *testing.T) {
	feature := feato.NewFeatureWithID("some-id")
	require.Equal(t, feato.DisableIndexToggle, feature.DefaultIndexToggle)
	require.Equal(t, "some-id", feature.ID)
}

func TestFeature_New(t *testing.T) {
	feature1 := feato.NewFeature()
	feature2 := feato.NewFeature()
	require.NotEqual(t, feature1.ID, feature2.ID)
}

func TestFeature_SetDefaultIndexToggle(t *testing.T) {
	feature := feato.NewFeature().SetDefaultIndexToggle(99)
	require.Equal(t, feato.IndexToggle(99), feature.DefaultIndexToggle)
}
