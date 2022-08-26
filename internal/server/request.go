package server

import (
	"context"
	"net/http"
)

func WithParsedBody(req *http.Request, body []byte) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "parsed_body", body)
	return req.WithContext(ctx)
}

func GetParsedBody(req *http.Request) ([]byte, bool) {
	body, ok := req.Context().Value("parsed_body").([]byte)
	return body, ok
}
