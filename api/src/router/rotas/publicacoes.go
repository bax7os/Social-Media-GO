package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		MetodoHttp:         http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI:                "/publicacoes",
		MetodoHttp:         http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: false,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		MetodoHttp:         http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		MetodoHttp:         http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
