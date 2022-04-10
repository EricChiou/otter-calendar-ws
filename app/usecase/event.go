package usecase

import (
	"otter-calendar/app/model/event"
	"otter-calendar/app/repository"
)

var Event = eventUsecase{}

type eventUsecase struct{}

func (u eventUsecase) AddEvent(e event.Entity) error {
	_, err := repository.Event.AddEvent(e)
	return err
}

func (u eventUsecase) GetEventListByUserID(userID int) ([]event.Entity, error) {
	eventList := []event.Entity{}

	result, err := repository.Event.GetEventListByUserID(userID)
	for result.Next() {
		var e event.Entity
		result.Scan(&e.ID, &e.Name, &e.Type, &e.StartTime, &e.RepeatUnit, &e.RepeatInterval, &e.RepeatTime, &e.LastTime, &e.Remark, &e.UserID)
		eventList = append(eventList, e)
	}

	return eventList, err
}
