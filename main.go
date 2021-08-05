package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", func(rw http.ResponseWriter, r *http.Request) {
		mySingingKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySingingKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})
	})

	api := r.NewRoute().Subrouter()
	api.Use(authMiddleware)

	r.HandleFunc("/todos", todo.AddTask).Methods(http.MethodPut).Subrouter()

	r.HandleFunc("/todos/{index}", todo.DoneTask).Methods(http.MethodPut)

	r.HandleFunc("/todos", todo.GetTask).Methods(http.MethodGet)

	http.ListenAndServe(":9090", r)

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer", "")
		mySingingKey := []byte("password")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySingingKey, nil
		})
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
