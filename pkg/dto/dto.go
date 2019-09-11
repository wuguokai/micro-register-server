package dto

import "time"

type Instance struct {
	ServiceName  string
	Ip           string
	RegisterDate time.Time
	RefreshDate  time.Time
	Status       string
}
