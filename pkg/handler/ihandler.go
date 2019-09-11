package handler

import "net/http"

type IHandler interface {
	Handle(res http.ResponseWriter, req *http.Request)
}
