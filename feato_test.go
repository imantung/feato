package feato_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/feato"
)

func TestManager(t *testing.T) {
	ctrl := gomock.NewController(t)
	toggleRouter := feato.NewMockToggleRouter(ctrl)
	feato.SetToggleRouter(toggleRouter)

	t.Run("GIVEN route return enable", func(t *testing.T) {
		toggleRouter.EXPECT().Route(gomock.Any()).Return(feato.EnableIndexToggle)
		err := feato.Run(feato.NewUniqueFeature(), func() error {
			return fmt.Errorf("error-but-enable")
		})
		require.EqualError(t, err, "error-but-enable")
	})

	t.Run("GIVEN route return disable", func(t *testing.T) {
		t.Run("WHEN no disable run function", func(t *testing.T) {
			toggleRouter.EXPECT().Route(gomock.Any()).Return(feato.DisableIndexToggle)
			err := feato.Run(feato.NewUniqueFeature(), func() error {
				return fmt.Errorf("error-but-enable")
			})
			require.EqualError(t, err, string(feato.ErrOutOfRunFunctionsIndex))
		})
		t.Run("When disable run function available", func(t *testing.T) {
			toggleRouter.EXPECT().Route(gomock.Any()).Return(feato.DisableIndexToggle)
			err := feato.Run(
				feato.NewUniqueFeature(),
				func() error {
					return fmt.Errorf("error-but-enable")
				},
				func() error {
					return fmt.Errorf("error-but-disable")
				},
			)
			require.EqualError(t, err, "error-but-disable")
		})

	})

}
