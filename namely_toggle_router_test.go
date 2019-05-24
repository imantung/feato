package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/feato"
)

func TestNamelyToggleRouter_ImplementationOfToggleRouter(t *testing.T) {
	var _ feato.ToggleRouter = (*feato.NamelyToggleRouter)(nil)
}

func TestNamelyToggleRouter_Route(t *testing.T) {
	toggleRouter := feato.NewNamelyToggleRouter()

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
		toggleRouter.SetIndexToggle("feature3", feato.EnableIndexToggle)
		require.Equal(t,
			feato.EnableIndexToggle,
			toggleRouter.Route(feato.NewFeature("feature3")))
	})

	t.Run("GIVEN feature is disable", func(t *testing.T) {
		toggleRouter.SetIndexToggle("feature4", feato.DisableIndexToggle)
		require.Equal(t,
			feato.DisableIndexToggle,
			toggleRouter.Route(feato.NewFeature("feature4")))
	})

}
