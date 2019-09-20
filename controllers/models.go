package controllers

import (
    "time"
    "html/template"
)

const API = "/api/v1"


type Page struct {
    Count int
    Posts  []Post
}

type Post struct {
    Id int
	Created time.Time
	Title string
	Text template.HTML
    Publish bool
    Language string
}
