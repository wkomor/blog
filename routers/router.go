package routers

import (
	"github.com/gorilla/mux"
	"komorovski/controllers"
)

func SetupRoutes() mux.Router {
	var Router = mux.NewRouter()
	Router.HandleFunc("/", controllers.MainPage)
	Router.HandleFunc("/products/{id:[0-9]+}", controllers.PostHandler)

	return *Router
}
