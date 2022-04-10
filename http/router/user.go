package router

import (
	"otter-calendar/app/delivery"
)

func initUserAPI() {
	groupName := "/user"

	post(groupName+"/signup", false, delivery.User.SignUp)
	post(groupName+"/login", false, delivery.User.Login)
}
