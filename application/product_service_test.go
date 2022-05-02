package application_test

// para criar mocks vamos utilizar o mockgen
// mockgen -destination=application/mocks/application.go -source=application/product.go

import (
	"testing"

	"github.com/GabrielBrotas/hexagonal-architeture/application"
	mock_application "github.com/GabrielBrotas/hexagonal-architeture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	// "github.com/GabrielBrotas/hexagonal-architeture/application"
	// "github.com/google/uuid"
)

func TestProductService_GetById(t *testing.T) {
	controller := gomock.NewController(t)

	// defer vai esperar tudo que esta acontecendo acabar para depois executar esse comando
	// posterga a execução
	defer controller.Finish()

	product := mock_application.NewMockIProduct(controller)
	persistence := mock_application.NewMockIProductPersistance(controller)

	// sempre que o metodo GetById for chamado vai retornar um mock product
	// simular que quando o nosso service for chamar o GetById vai retornar um produto,
	// pois nao queremos testar o externo, apenas o nosso servico
	persistence.EXPECT().GetById(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistance: persistence,
	}

	result, err := service.GetById("abc")

	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductServce_Create(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	product := mock_application.NewMockIProduct(controller)
	persistence := mock_application.NewMockIProductPersistance(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistance: persistence,
	}

	result, err := service.Create("Product 1", 10.00)

	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductServce_EnableDisable(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	product := mock_application.NewMockIProduct(controller)
	product.EXPECT().Enable().Return(nil) // nao retornar erro
	product.EXPECT().Disable().Return(nil) // nao retornar erro

	persistence := mock_application.NewMockIProductPersistance(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistance: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}