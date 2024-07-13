package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 500, "An error occurred")
}
