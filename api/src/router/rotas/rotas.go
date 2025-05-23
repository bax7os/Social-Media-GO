package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota represents all API routes
type Rota struct {
	URI                string
	MetodoHttp         string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar defines all API routes
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.MetodoHttp)
	}

	return r
}
