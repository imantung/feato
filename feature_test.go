package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/feato"
)

func TestFeature_New(t *testing.T) {
	feature := feato.NewFeature("some-feature")
	require.Equal(t, feato.DisableIndexToggle, feature.DefaultIndexToggle)
}

func TestFeature_SetDefaultIndexToggle(t *testing.T) {
	feature := feato.NewFeature("some-feature").SetDefaultIndexToggle(99)
	require.Equal(t, feato.IndexToggle(99), feature.DefaultIndexToggle)
}
