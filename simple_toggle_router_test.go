package feato_test

import (
	"testing"

	"github.com/imantung/feato"
	"github.com/stretchr/testify/require"
)

func TestSimpleToggleRouter_ImplementationOfToggleRouter(t *testing.T) {
	var _ feato.Router = (*feato.SimpleToggleRouter)(nil)
}

func TestSimpleToggleRouter_Route(t *testing.T) {
	toggleRouter := feato.NewSimpleToggleRouter()

	t.Run("GIVEN feature map is not set", func(t *testing.T) {
		t.Run("WHEN default index toggle is disable", func(t *testing.T) {
			require.Equal(t,
				feato.Disabled,
				toggleRouter.Route(feato.NewFeature("key1")))
		})
		t.Run("WHEN default index toggle is enable", func(t *testing.T) {
			require.Equal(t,
				feato.Enabled,
				toggleRouter.Route(feato.NewFeature("key2").EnableByDefault()))
		})
	})

	t.Run("GIVEN feature is enable", func(t *testing.T) {
		feature3 := feato.NewFeature("key3")
		toggleRouter.Enable(feature3)
		require.Equal(t, feato.Enabled, toggleRouter.Route(feature3))
	})

	t.Run("GIVEN feature is disable", func(t *testing.T) {
		feature4 := feato.NewFeature("key4")
		toggleRouter.Disable(feature4)
		require.Equal(t, feato.Disabled, toggleRouter.Route(feature4))
	})
}
