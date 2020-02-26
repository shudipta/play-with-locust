package mock

import (
	"fmt"
	"net/http"
)

type Response struct {
	statusCode int
	msg        string
}

func Respond(w http.ResponseWriter, r Response) {
	if r.statusCode == http.StatusUnauthorized {
		w.Header().Add("WWW-Authenticate", `Basic realm="Authorization Required"`)
	}
	w.WriteHeader(r.statusCode)
	fmt.Fprintf(w, r.msg)
}

func NewResponse(statusCode int, msg string) Response {
	return Response{statusCode, msg}
}
