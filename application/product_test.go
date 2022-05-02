package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/GabrielBrotas/hexagonal-architeture/application"
	"github.com/google/uuid"
)

// testar a struct com o metodo Enable
func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	// o erro tem que ser nil
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Product price must be greater than 0", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Product price must be equal to 0", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID: uuid.New().String(),
		Name: "Product 1",
		Status: application.DISABLED,
		Price: 10,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "invalid status", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "product price must be greater than 0", err.Error())
}