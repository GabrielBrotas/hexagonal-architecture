package cli_test

import (
	"fmt"
	"testing"

	"github.com/GabrielBrotas/hexagonal-architeture/adapters/cli"
	mock_application "github.com/GabrielBrotas/hexagonal-architeture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productName := "Product 1"
	productPrice := 20.99
	productStatus := "enabled"

	productMock := mock_application.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockIProductService(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().GetById(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultCreateExpected := fmt.Sprintf(
		"Product ID %s with Name %s has been created with price %f and status %s", 
		productId, 
		productName, 
		productPrice,
		productStatus,
	)

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultCreateExpected, result)

	resultEnableExpected := fmt.Sprintf(
		"Product ID %s with Name %s has been Enabled", 
		productId, 
		productName, 
	)

	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultEnableExpected, result)

	resultDisableExpected := fmt.Sprintf(
		"Product ID %s with Name %s has been Disabled", 
		productId, 
		productName, 
	)

	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultDisableExpected, result)

	resultGetProductExpected := fmt.Sprintf(
		"Product ID: %s \nName: %s\nPrice: %f\nStatus: %s", 
		productId, 
		productName,
		productPrice,
		productStatus, 
	)

	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultGetProductExpected, result)
}