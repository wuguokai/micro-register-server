package cmd

import (
	"micro-register-server/pkg/config"
	"micro-register-server/pkg/util"
	"net/http"
)

func ServerStart() {
	config.Config()
	http.HandleFunc("/", config.ReqHandler)
	err := http.ListenAndServe(":8000", nil)
	util.CheckErr(err)
}
