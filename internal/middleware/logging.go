package middleware

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AleksandrCherepanov/tg-bot/internal/server"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Request. %s %s %d %s", req.Method, req.RequestURI, start.UnixMilli(), err.Error())
		}

		log.Printf("Request. %s %s %d %s", req.Method, req.RequestURI, start.UnixMilli(), string(body))
		req = server.WithParsedBody(req, body)
		next.ServeHTTP(w, req)
		log.Printf("Response. %s %s %d", req.Method, req.RequestURI, time.Since(start).Milliseconds())
	})
}
