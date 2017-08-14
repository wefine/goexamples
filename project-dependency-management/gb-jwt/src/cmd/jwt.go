package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/codegangsta/negroni"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
)

var SingedKey = []byte("wefine is really fine")

func fatal(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type UserCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Response struct {
    Data string `json:"data"`
}

type Token struct {
    Token string `json:"token"`
}

func StartServer() {

    http.HandleFunc("/login", LoginHandler)

    http.Handle("/resource", negroni.New(
        negroni.HandlerFunc(ValidateTokenMiddleware),
        negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
    ))

    port := ":5000"
    log.Println("Now listening..." + port)
    http.ListenAndServe(port, nil)
}

func main() {
    StartServer()
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

    response := Response{"Gained access to protected resource"}
    JsonResponse(response, w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

    var user UserCredentials

    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        w.WriteHeader(http.StatusForbidden)
        fmt.Fprint(w, "Error in request")
        return
    }

    if strings.ToLower(user.Username) != "someone" {
        if user.Password != "p@ssword" {
            w.WriteHeader(http.StatusForbidden)
            fmt.Println("Error logging in")
            fmt.Fprint(w, "Invalid credentials")
            return
        }
    }

    claims := &jwt.StandardClaims{
        // 过期时间，10小时
        ExpiresAt: time.Now().Add(time.Hour * time.Duration(10)).Unix(),
        // 签发时间
        IssuedAt: time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(SingedKey)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintln(w, "Error while signing the token")
        fatal(err)
    }

    response := Token{tokenString}
    JsonResponse(response, w)
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

    token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
        func(token *jwt.Token) (interface{}, error) {
            return SingedKey, nil
        })

    if err == nil {
        if token.Valid {
            next(w, r)
        } else {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Token is not valid")
        }
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprint(w, "Unauthorized access to this resource")
    }
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

    bytes, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write(bytes)
}
