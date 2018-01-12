package main

import (
    "encoding/json"
    "log"
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
    "net/url"
)

var templates *template.Template 
type user struct {
    Name, Email, Emoji string
    Colors []string
}
type token struct {
    AccessToken string `json:"access_token"`
}

func login(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "GET":
            fmt.Println("Doing a GET")
            error := templates.ExecuteTemplate(res, "login.gohtml", nil)
            if error != nil {
                log.Fatalln(error)
            }
        case "POST":
            fmt.Println("Doing a POST")
            email := req.FormValue("email")
            password := req.FormValue("password") 

            formData := url.Values{}
            formData.Add("email", email)
            formData.Add("password", password)

            t := token{}
            resp, error := http.PostForm("http://127.0.0.1:5000/auth/token", formData)
            defer resp.Body.Close()
            error = json.NewDecoder(resp.Body).Decode(&t)
            fmt.Println(t)


            data := struct {
                Email, Password, Token string 
            }{
                Email: email,
                Password: password,
                Token: t.AccessToken,
            }

            error = templates.ExecuteTemplate(res, "home.gohtml", data)  
            if error != nil {
                log.Fatalln(error)
                fmt.Println("There was an error on the POST")
            }
        default:
            fmt.Println("IDK what I'm doing :D")
    }
}

func greetUser(res http.ResponseWriter, req *http.Request) {
    newReq, error := 
    resp, error := http.Get("http://127.0.0.1:5000/users/greet")
    if error != nil {
        log.Fatalln(error)
    }

    body, error := ioutil.ReadAll(resp.Body)
    if error != nil {
        log.Fatalln(error)
    }
    fmt.Println(string(body))

    b := user {
        Name: "Balloonicorn",
        Email: "b@sparkles.com",
        Emoji: "💖",
        Colors: []string{"pink", "blue", "rainbow"},
    }

    templates.ExecuteTemplate(res, "greeting.gohtml", b)
}

func init() {
    templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    fmt.Println("Now serving...")
    http.HandleFunc("/login", login)
    http.HandleFunc("/greet", greetUser)

    // When handler is nil, we're using Go's DefaultServeMux.
    http.ListenAndServe(":8080", nil)
}