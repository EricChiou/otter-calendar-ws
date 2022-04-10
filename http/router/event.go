package router

import (
	"otter-calendar/app/delivery"
)

func initEventAPI() {
	groupName := "/event"

	post(groupName+"/add", true, delivery.Event.AddEvent)
	get(groupName, true, delivery.Event.GetEventListByUserID)
}
