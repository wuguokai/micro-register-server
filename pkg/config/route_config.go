package config

import (
	"micro-register-server/pkg/handler"
	"net/http"
)

var pathConfig = make(map[string]handler.IHandler)

func Config() {
	pathConfig["POST - /register"] = handler.RegisterHandler{}
	pathConfig["GET - /query"] = handler.QueryHandler{}
}

func ReqHandler(res http.ResponseWriter, req *http.Request) {
	requestPath := req.URL.Path
	requestMethod := req.Method
	requestKey := requestMethod + " - " + requestPath
	for k, v := range pathConfig {
		if requestKey == k {
			v.Handle(res, req)
		}
	}
}
