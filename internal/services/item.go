package services

import (
	"encoding/json"
	"errors"
	"myapi/internal/models"
	"myapi/validators"
	"net/http"
)

// service para criar um item
func DecodeAndValidate(w http.ResponseWriter, r *http.Request) (*models.Item, error) {
	var item models.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		return nil, errors.New("Erro ao decodificar o item")
	}

	if err := validators.ValidateItem(&item); err != nil {
		return nil, errors.New("Erro de Validação" + err.Error())

	}

	return &item, nil
}
