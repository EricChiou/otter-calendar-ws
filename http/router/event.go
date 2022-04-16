package router

import (
	"otter-calendar/app/delivery"
	"otter-calendar/config"
)

func initEventAPI() {
	groupName := "/event"
	if config.Get().ENV == "prod" {
		groupName = "/otter-calendar-ws" + groupName
	}

	post(groupName+"/add", true, delivery.Event.AddEvent)
	get(groupName, true, delivery.Event.GetEventListByUserID)
	get(groupName+"/:id", true, delivery.Event.GetEventByEventID)
	put(groupName, true, delivery.Event.UpdateEvent)
	delete(groupName+"/:id", true, delivery.Event.DeleteEventByEventID)
}
