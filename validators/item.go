package validators

import (
	"errors"
	"myapi/internal/models"
)

func ValidateItem(item *models.Item) error {
	if item.Preco <= 0 {
		return errors.New("o preço deve ser maior que zero")
	}

	if item.Quantidade < 0 {
		return errors.New("a quantidade deve ser igual ou maior que zero")
	}

	if len(item.Codigo) != 6 {
		return errors.New("o código precisa ter 6 caracteres")
	}

	return nil
}
