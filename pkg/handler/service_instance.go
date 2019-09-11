package handler

import (
	"io"
	"micro-register-server/pkg/dto"
	"micro-register-server/pkg/util"
	"net/http"
	"time"
)

var instances = make(map[string][]dto.Instance)

func Register(instance dto.Instance, res http.ResponseWriter) {
	if !checkExist(instance) {
		//服务未注册，添加实例
		instance.RegisterDate = time.Now()
		instance.RefreshDate = time.Now()
		currentInstances := append(instances[instance.ServiceName], instance)
		instances[instance.ServiceName] = currentInstances
		_, err := io.WriteString(res, "注册成功")
		util.CheckErr(err)
	} else {
		//服务已注册,修改刷新时间
		instanceArrCurrent := instances[instance.ServiceName]
		for k, v := range instanceArrCurrent {
			if v.Ip == instance.Ip {
				v.RefreshDate = time.Now()
				instanceArrCurrent[k] = v
			}
		}
		instances[instance.ServiceName] = instanceArrCurrent
		_, err := io.WriteString(res, "实例已注册，刷新成功")
		util.CheckErr(err)
	}
}

func Query() map[string][]dto.Instance {
	return instances
}

func checkExist(instance dto.Instance) bool {
	for k, instancesArr := range instances {
		if k == instance.ServiceName {
			for _, i := range instancesArr {
				if i.Ip == instance.Ip {
					return true
				}
			}
		}
	}
	return false
}
