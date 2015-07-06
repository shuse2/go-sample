package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"time"
)

type appContext struct {
	db *sql.DB
}

func (c *appContext) authHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		authToken := req.Header.Get("Authentication")
		user, err := getUser(authToken)

		if err != nil {
			http.Error(res, http.StatusText(401), 401)
			return
		}

		context.Set(req, "user", user)
		next.ServeHTTP(res, req)
	}
	return http.HandlerFunc(fn)
}

func (c *appContext) adminHandler(res http.ResponseWriter, req *http.Request) {
	user := context.Get(req, "user")
	json.NewEncoder(res).Encode(user)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(res, req)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", req.Method, req.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func recoveryHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %+v", err)
				http.Error(res, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index handler")
}

func getUser(token string) (user string, err error) {
	return user, err
}

func main() {
	db, err := sql.Open("mysql", "err")
	if err != nil {
		log.Println("supposed to show error")
	}
	appC := appContext{db}
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoveryHandler, appC.authHandler)
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.ListenAndServe(":8080", nil)
}
