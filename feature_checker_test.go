package feato_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typfeat"
)

func TestFeatureChecker_IsEnabled(t *testing.T) {
	checker := typfeat.NewFeatureChecker()

	t.Run("Given no listed feature", func(t *testing.T) {
		checker.Reset()
		t.Run("When DefaultState is false", func(t *testing.T) {
			require.False(t, checker.IsEnabled(typfeat.Feature{
				Name:         "feature01",
				DefaultState: false,
			}))
		})
		t.Run("When DefaultState is true", func(t *testing.T) {
			require.True(t, checker.IsEnabled(typfeat.Feature{
				Name:         "feature02",
				DefaultState: true,
			}))
		})
	})

	t.Run("Given listed feature", func(t *testing.T) {
		t.Run("When listed feature state is false", func(t *testing.T) {
			checker.Put("feature03", true)
			require.True(t, checker.IsEnabled(typfeat.Feature{
				Name:         "feature03",
				DefaultState: false,
			}))
		})
		t.Run("When listed feature state is true", func(t *testing.T) {
			checker.Put("feature04", false)
			require.False(t, checker.IsEnabled(typfeat.Feature{
				Name:         "feature04",
				DefaultState: true,
			}))
		})
	})
}

func TestFeatureCheck_Put(t *testing.T) {
	var feature *typfeat.Feature
	name := "same-name"

	checker := typfeat.NewFeatureChecker()
	checker.Put(name, true)

	feature = checker.Get(name)
	require.Equal(t, &typfeat.Feature{
		Name:         name,
		DefaultState: false,
		State:        true,
	}, feature)

	checker.Put(name, false)
	require.False(t, feature.State)

	checker.Delete(name)
	require.Nil(t, checker.Get(name))
}
