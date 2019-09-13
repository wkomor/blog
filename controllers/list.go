package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/biezhi/gorm-paginator/pagination"
	"strconv"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	db := connectToDB()
	var posts []Post
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
    if err != nil {
		page = 1
	} 
	limit := 5
	
	res := db.Where("publish = ?", true)
	pagination.Paging(&pagination.Param{
		DB:      res,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created desc"},
	}, &posts)
	defer db.Close()
	servePostsJson(w, posts)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}
