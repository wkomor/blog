package routers

import (
	"github.com/gorilla/mux"
	"komorovski/controllers"
)


func SetupRoutes() mux.Router {
	var Router = mux.NewRouter()
	Router.HandleFunc(controllers.API + "/", controllers.MainPage)
	Router.HandleFunc(controllers.API + "/post/{id:[0-9]+}", controllers.PostHandler)

	return *Router
}
