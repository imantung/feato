package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/feato"
)

func TestNamelyToggleRouter_ImplementOfToggleRouter(t *testing.T) {
	var _ feato.ToggleRouter = (*feato.NamelyToggleRouter)(nil)
}

func TestNamelyToggleRouter_Route(t *testing.T) {
	toggleRouter := feato.NewNamelyToggleRouter()

	t.Run("WHEN feature map is not set", func(t *testing.T) {
		require.Equal(t, feato.DisableIndexToggle, toggleRouter.Route(feato.Feature{Name: "feature1", DefaultState: feato.DisableIndexToggle}))
		require.Equal(t, feato.EnableIndexToggle, toggleRouter.Route(feato.Feature{Name: "feature1", DefaultState: feato.EnableIndexToggle}))
	})

	t.Run("WHEN feature is enable", func(t *testing.T) {
		toggleRouter.SetState("feature2", feato.EnableIndexToggle)
		require.Equal(t, feato.EnableIndexToggle, toggleRouter.Route(feato.Feature{Name: "feature2"}))
	})

	t.Run("WHEN feature is disable", func(t *testing.T) {
		toggleRouter.SetState("feature3", feato.DisableIndexToggle)
		require.Equal(t, feato.DisableIndexToggle, toggleRouter.Route(feato.Feature{Name: "feature3"}))
	})

}
