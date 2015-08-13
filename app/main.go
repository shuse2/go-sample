package main

import (
	"flag"
	"github.com/golang/glog"
	// "errors"
	"github.com/go-sample/app/controllers"
	"github.com/go-sample/app/core"
	"github.com/go-sample/app/loggers"
	"github.com/go-sample/app/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

// TODO: what should i do with this....
func wrapHandler(h http.Handler) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		context.Set(req, "params", params)
		h.ServeHTTP(res, req)
	}
}

func main() {
	filename := flag.String("config", "config/local.json", "Path to configuration file")
	flag.Parse()

	defer glog.Flush()

	var application = &core.Application{}
	application.Init(*filename)
	models.Init(application.Configuration)
	defer models.Close()

	// set up middleware
	commonHandlers := alice.New(context.ClearHandler, loggers.LoggingHandler, application.RecoveryHandler)

	// setup routes
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.GET("/", wrapHandler(commonHandlers.ThenFunc(controllers.IndexHandler)))
	commonHandlers = commonHandlers.Append(application.ContextTypeHandler)
	router.GET("/twitter/login", wrapHandler(commonHandlers.ThenFunc(controllers.TwitterLoginHandler)))
	router.GET("/twitter/callback", wrapHandler(commonHandlers.ThenFunc(controllers.TwitterLoginCallbackHandler)))
	// add auth handler
	commonHandlers = commonHandlers.Append(application.AuthHandler)
	router.GET("/api/user", wrapHandler(commonHandlers.ThenFunc(controllers.UserHandler)))

	glog.Info("Starting server at 3000")
	http.ListenAndServe(":3000", router)
}
