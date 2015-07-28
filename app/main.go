package main

import (
	"flag"
	"github.com/golang/glog"
	// "errors"
	"github.com/go-sample/app/controllers"
	"github.com/go-sample/app/core"
	"github.com/go-sample/app/loggers"
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

	// set up middleware
	commonHandlers := alice.New(context.ClearHandler, loggers.LoggingHandler, application.RecoveryHandler, application.AuthHandler)

	// setup routes
	router := httprouter.New()
	router.GET("/", wrapHandler(commonHandlers.ThenFunc(controllers.IndexHandler)))

	http.ListenAndServe(":8080", router)
}
