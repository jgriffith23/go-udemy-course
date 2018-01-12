package main

import (
    "html/template"
    "log"
    "net/http"
)

// FIXME: Use conventiona/ type
type server string

// We'll want a pointer to a Template object to be able to execute our 
// templates later.
var templates *template.Template

// Add a value method to server to make it able to serve files.
func (s server) ServeHTTP (resp http.ResponseWriter, req *http.Request) {
    error := req.ParseForm()
    if error != nil {
        log.Fatalln(error)
    }
    
    // Render a specific template to the response.
    templates.ExecuteTemplate(resp, "login.gohtml", nil)
}

// For convenience, load templates in the init().

func init() {
    // templates will now point to a Template glob that has all templates.
    // Let Must function handle our errors for us.
    templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    var srv server
    http.ListenAndServe(":8000", srv)
}