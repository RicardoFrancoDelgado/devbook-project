package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao adiciona uma nova publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairTokenID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao modelos.Publicacao

	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioID
	
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusCreated, publicacao)

}

// BuscarPublicacoes traz as  publicacões que apareceriam no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao traz uma unica publicacao
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao altera os dados de uma unica publicacao
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// DeletarPublicacao exclui os dados de uma publicacao
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
