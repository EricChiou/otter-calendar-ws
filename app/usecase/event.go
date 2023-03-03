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
		result.Scan(&e.ID, &e.Name, &e.Type, &e.StartTime, &e.RepeatUnit, &e.RepeatInterval, &e.RepeatTime, &e.LastTime, &e.Remark, &e.UserID, &e.CalType)
		eventList = append(eventList, e)
	}

	return eventList, err
}

func (u eventUsecase) GetEventByEventID(eventID, userID int) (event.Entity, error) {
	var e event.Entity

	result := repository.Event.GetEventByEventID(eventID, userID)
	err := result.Scan(&e.ID, &e.Name, &e.Type, &e.StartTime, &e.RepeatUnit, &e.RepeatInterval, &e.RepeatTime, &e.LastTime, &e.Remark, &e.UserID, &e.CalType)

	return e, err
}

func (u eventUsecase) UpdateEvent(e event.Entity) error {
	_, err := repository.Event.UpdateEvent(e)
	return err
}

func (u eventUsecase) DeleteEventByEventID(eventID, userID int) error {
	_, err := repository.Event.DeleteEventByEventID(eventID, userID)
	return err
}
