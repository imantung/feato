package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/feato"
)

func TestSimpleToggleRouter_ImplementationOfToggleRouter(t *testing.T) {
	var _ feato.ToggleRouter = (*feato.SimpleToggleRouter)(nil)
}

func TestSimpleToggleRouter_Route(t *testing.T) {
	toggleRouter := feato.NewSimpleToggleRouter()

	t.Run("GIVEN feature map is not set", func(t *testing.T) {
		t.Run("WHEN default index toggle is disable", func(t *testing.T) {
			require.Equal(t,
				feato.DisableIndexToggle,
				toggleRouter.Route(&feato.Feature{DefaultIndexToggle: feato.DisableIndexToggle}))
		})
		t.Run("WHEN default index toggle is enable", func(t *testing.T) {
			require.Equal(t,
				feato.EnableIndexToggle,
				toggleRouter.Route(&feato.Feature{DefaultIndexToggle: feato.EnableIndexToggle}))
		})
	})

	t.Run("GIVEN feature is enable", func(t *testing.T) {
		feature3 := feato.NewUniqueFeature()
		toggleRouter.Enable(feature3)
		require.Equal(t, feato.EnableIndexToggle, toggleRouter.Route(feature3))
	})

	t.Run("GIVEN feature is disable", func(t *testing.T) {
		feature4 := feato.NewUniqueFeature()
		toggleRouter.Disable(feature4)
		require.Equal(t, feato.DisableIndexToggle, toggleRouter.Route(feature4))
	})

}
