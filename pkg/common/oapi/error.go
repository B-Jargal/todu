package oapi

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Shortcut that sends error response immediately
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}

func Forbidden(w http.ResponseWriter) {
	ClientError(w, http.StatusForbidden)
}
