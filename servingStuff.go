package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "net/url"
)

type server int
var templates *template.Template 

func (s server) ServeHTTP (w http.ResponseWriter, req *http.Request) {
    error := req.ParseForm()
    if error != nil {
        log.Fatalln(error)
    }

    // Have the response writer set the authorization header. This is where
    // a JWT might go.
    w.Header().Set("Authorization", "Bearer SOME JWT THING")

    data := struct {
            Method        string
            URL           *url.URL
            Submissions   url.Values
            Header        http.Header
            Host          string
            ContentLength int64
        }{
            req.Method,
            req.URL,
            req.Form,
            req.Header,
            req.Host,
            req.ContentLength,
        }

    templates.ExecuteTemplate(w, "login2.gohtml", data)
}

func init() {
    templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    var srv server
    fmt.Println("Now serving...")
    http.ListenAndServe(":8080", srv)
}