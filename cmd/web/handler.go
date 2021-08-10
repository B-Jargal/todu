package main

import (
	"net/http"
	// "git.bolor.net/bolorsoft/qa/pkg/application"
	// "git.bolor.net/bolorsoft/qa/pkg/common/oapi"
	// "git.bolor.net/bolorsoft/qa/pkg/common/ocookie"
)

const (
	MB = 1 << 20
	KB = 1 << 10
)

func cacheUserID(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := app.Blogin.CacheAuthToken(r); err != nil {
			oapi.ClientError(w, http.StatusBadRequest)
			return
		}
		w.Write([]byte("OK"))
	}
}


func 
