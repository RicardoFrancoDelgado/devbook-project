package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroMsg struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, ErroMsg{
		Erro: erro.Error(),
	})
}
