package router

import (
	"otter-calendar/app/delivery"
)

func initEventAPI() {
	groupName := "/event"

	post(groupName+"/add", true, delivery.Event.AddEvent)
	get(groupName, true, delivery.Event.GetEventListByUserID)
	get(groupName+"/:id", true, delivery.Event.GetEventByEventID)
	put(groupName, true, delivery.Event.UpdateEvent)
	delete(groupName+"/:id", true, delivery.Event.DeleteEventByEventID)
}
