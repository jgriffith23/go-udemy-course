package main 

import (
    "fmt"
    "net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
    fmt.Println("home")
}

func dog(res http.ResponseWriter, req *http.Request) {
    fmt.Println("dog")
}

func me(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Your Name")
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/dog/", dog)
    http.HandleFunc("/me/", me)

    http.ListenAndServe(":8080", nil)
}