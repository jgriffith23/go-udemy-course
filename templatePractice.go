package main

import (
    "log"
    "os"
    "html/template"
)

type user struct {
    Name, Email, Emoji string
    Colors []string
}

// We'll want a pointer to a Template object to be able to execute our 
// templates later.
var templates *template.Template

// For convenience, load templates in the init().

func init() {
    // templates will now point to a Template glob that has all templates.
    // Let Must function handle our errors for us.
    templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

    b := user {
        Name: "Balloonicorn",
        Email: "b@sparkles.com",
        Emoji: "💖",
        Colors: []string{"pink", "blue", "rainbow"},
    }

    // Render a specific template to stdout. 
    error := templates.ExecuteTemplate(os.Stdout, "greeting.gohtml", b)

    if error != nil {
        log.Fatalln(error)
    }

}