package router

import (
	"otter-calendar/app/delivery"
	"otter-calendar/config"
)

func initUserAPI() {
	groupName := "/user"
	if config.Get().ENV == "prod" {
		groupName = "/otter-calendar-ws" + groupName
	}

	post(groupName+"/signup", false, delivery.User.SignUp)
	post(groupName+"/login", false, delivery.User.Login)
}
