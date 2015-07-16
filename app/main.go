package main

import (
	"database/sql"
	"encoding/json"
	// "errors"
	"fmt"
	"github.com/go-sample/app/loggers"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"log"
	"net/http"
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
	// err = errors.New("this is error")
	return user, err
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		context.Set(req, "params", params)
		h.ServeHTTP(res, req)
	}
}

func main() {
	db, err := sql.Open("mysql", "err")
	if err != nil {
		log.Println("supposed to show error")
	}
	appC := appContext{db}
	commonHandlers := alice.New(context.ClearHandler, loggers.LoggingHandler, recoveryHandler, appC.authHandler)
	router := httprouter.New()
	router.GET("/", wrapHandler(commonHandlers.ThenFunc(indexHandler)))
	http.ListenAndServe(":8080", router)
}
