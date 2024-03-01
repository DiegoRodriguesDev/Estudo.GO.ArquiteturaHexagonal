package application_test

import (
	"testing"

	"github.com/DiegoRodriguesDev/Estudo.GO.ArquiteturaHexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "Price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)

	product.Status = "INVALID"

	isValid, err = product.IsValid()
	require.Equal(t, "Status must be enabled or disabled", err.Error())
	require.Equal(t, false, isValid)

	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)

	product.Price = -10
	isValid, err = product.IsValid()
	require.Equal(t, "Price must be greater or equal zero", err.Error())
	require.Equal(t, false, isValid)
}
