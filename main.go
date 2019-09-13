package main

import (
	"fmt"
	"net/http"
	"komorovski/routers"
)

func main() {
	router := routers.SetupRoutes()
	/* http.Handle("/", &router) */
	fmt.Println("Server is listening...")
	
	http.ListenAndServe(":8181",  &router)
}
