package application_test

import (
	"testing"

	"github.com/DiegoRodriguesDev/Estudo.GO.ArquiteturaHexagonal/application"
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
