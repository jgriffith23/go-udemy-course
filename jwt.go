package main

import (
    "fmt"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "time"
    "os"
    "encoding/json"
)

func generateJWT(subject string, mult time.Duration) string {
    // Get JWT secret out of environment variable
    secret := []byte(os.Getenv("JWT_SECRET"))

    // Put together some typically expected claims
    claims := &jwt.StandardClaims { 
        ExpiresAt: time.Now().Add(time.Second * 1800).Unix(),
        Subject: subject,
    }

    // Create the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token string with our secret
    signedStr, err := token.SignedString(secret)

    if err != nil {
        fmt.Println(err)
    }

    return signedStr
}

func validateJWT(signedToken string) string {
    // Use the secret from our environment to parse the token contents.
    token, err := jwt.ParseWithClaims(signedToken, &jwt.StandardClaims{}, func (token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    // This error will catch whether the token was expired.
    if err != nil {
        return "expired"
    }

    method := token.Method.Alg()

    // The token is ok if it parsed without error and uses the correct method. 
    if token.Valid && method == "HS256" {
        // Cast the token's Claims into StandardClaims so we can extract the payload
        subject := token.Claims.(*jwt.StandardClaims).Subject

        return "Hello, " +  subject + "!"
    }
    return "invalid"
}

func authenticateUser(res http.ResponseWriter, req *http.Request) {
    // Maybe this should talk to the database.
    name := req.FormValue("name")
    signedToken := generateJWT(name, 1800)

    res.Header().Set("Content-Type", "application/json")

    token := struct {
        AccessToken string `json:"access_token"`
    }{
        signedToken,
    }

    json.NewEncoder(res).Encode(token)
}

func getGreeting(res http.ResponseWriter, req *http.Request) {
    token := req.Header.Get("Authorization")

    greeting := struct {
        Message string `json:"message"`
    }{
        validateJWT(token),
    }

    json.NewEncoder(res).Encode(greeting)
}

func main() {

    // One route to generate the token and one to validate/make greeting
    http.HandleFunc("/auth/token", authenticateUser)
    http.HandleFunc("/greet", getGreeting)

    // When handler is nil, we're using Go's DefaultServeMux.
    http.ListenAndServe(":8080", nil)
}