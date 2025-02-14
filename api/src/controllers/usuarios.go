package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorios := repositorios.NovoRepositorioDeUsuario(db)
	IDInserido, erro := repositorios.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", IDInserido)))
}

func BuscandoUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando os usuários"))
}

func BuscandoUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

func DeletandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
