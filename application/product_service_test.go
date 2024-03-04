package application_test

import (
	"testing"

	"github.com/DiegoRodriguesDev/Estudo.GO.ArquiteturaHexagonal/application"
	mock_application "github.com/DiegoRodriguesDev/Estudo.GO.ArquiteturaHexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {

	controller := gomock.NewController(t)

	// o "defer" define que o comando controller.Finish() será executado apenas no final do método
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)

	// todas as vezes que o método Get() for chamado, será retornado o product de mock
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
