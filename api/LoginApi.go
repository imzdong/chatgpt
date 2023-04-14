package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var (
	jwtSecret = []byte("your_jwt_secret_key")
	users     = map[string]string{
		"admin": "admin",
		"test":  "test",
	}
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	credentials := Credentials{
		Username: user,
		Password: pass,
	}

	doAuth(w, r, &credentials)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	doAuth(w, r, &credentials)

}

func doAuth(w http.ResponseWriter, r *http.Request, credentials *Credentials) {
	if password, ok := users[credentials.Username]; ok && password == credentials.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			},
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "Hello, %s!", claims.Username)
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
}
