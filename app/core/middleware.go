package core

import (
	"github.com/go-sample/app/models"
	"github.com/golang/glog"
	"github.com/gorilla/context"
	"net/http"
)

func (application *Application) AuthHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		authToken := req.Header.Get("Authentication")
		user, err := models.GetUser("123", authToken)

		if err != nil {
			http.Error(res, http.StatusText(401), 401)
			return
		}

		context.Set(req, "user", user)
		next.ServeHTTP(res, req)
	}
	return http.HandlerFunc(fn)
}

func (application *Application) ContextTypeHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/vnd.api+json")
		next.ServeHTTP(res, req)
	}
	return http.HandlerFunc(fn)
}

// middleware?
func (application *Application) RecoveryHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				glog.Warningf("Panic: %+v", err)
				http.Error(res, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}
