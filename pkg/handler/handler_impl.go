package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"micro-register-server/pkg/dto"
	"micro-register-server/pkg/util"
	"net/http"
)

type RegisterHandler struct {
}

type QueryHandler struct {
}

func (handler RegisterHandler) Handle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("register")

	instance := dto.Instance{}
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(body))
	err := json.Unmarshal(body, &instance)
	util.CheckErr(err)
	instance.Status = "UP"

	fmt.Println(instance)
	if instance.ServiceName != "" {
		Register(instance, res)
	}
}

func (handler QueryHandler) Handle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("query")

	instances, _ := json.Marshal(Query())
	_, err := io.WriteString(res, string(instances))
	util.CheckErr(err)
}
