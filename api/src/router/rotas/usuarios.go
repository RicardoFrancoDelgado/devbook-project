package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscandoUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscandoUsuarioPorId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletandoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.Seguir,
		RequerAutenticacao: true,
	},
}
