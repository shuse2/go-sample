package loggers

import (
	"log"
	"net/http"
	"time"
)

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(res, req)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", req.Method, req.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}
