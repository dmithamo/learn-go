package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func readParamByName(name string, r *http.Request)string{
	return httprouter.ParamsFromContext(r.Context()).ByName(name)
}
