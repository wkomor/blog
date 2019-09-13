package controllers

import (
    "os"
    "net/http"
    "time"
    "log"
    "html/template"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // driver for mysql
)

type Page struct {
    Title string
    Posts  []Post
}

type Post struct {
	Created time.Time
	Title string
	Text template.HTML
	Publish bool
}


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl)
    t.Execute(w, p)
}

func servePostsJson(w http.ResponseWriter, posts []Post) {
    js, err := json.Marshal(posts)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(js)

}


func connectToDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    dbPass := os.Getenv("DB_PASS")
    dbUser := os.Getenv("DB_USER")
    connection := dbUser + ":" + dbPass + "@/komorovski?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open("mysql", connection)
    if err != nil {
      panic(err.Error() + " failed to connect database")
    }
    return db

  }