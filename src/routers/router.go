package routers

import (
	"api/src/routers/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
