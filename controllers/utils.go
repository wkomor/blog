package controllers

import (
    "os"
    "net/http"
    "log"
    "html/template"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"  // driver for mysql
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl)
    t.Execute(w, p)
}

func servePostsJson(w http.ResponseWriter, posts []Post) {
    var data Page
    data.Posts = posts
    data.Count = 22
    js, err := json.Marshal(data)
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
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    connection := "user=" + dbUser + " password=" + dbPass + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName
    db, err := gorm.Open("mysql", connection)
    if err != nil {
      panic(err.Error() + " failed to connect database")
    }
    return db
}