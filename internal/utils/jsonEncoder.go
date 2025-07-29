package utils

import (
	"encoding/json"
	"net/http"
)

func JsonEncode[T any](w http.ResponseWriter, item T) {
	if err := json.NewEncoder(w).Encode(item); err != nil {
		RespondWithError(w, "Erro ao codificar json", http.StatusInternalServerError)
		return
	}
}
