package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		MetodoHttp:         http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		MetodoHttp:         http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
