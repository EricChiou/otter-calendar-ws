package repository

import (
	"database/sql"
	"otter-calendar/app/model/event"
	"otter-calendar/db"
	"otter-calendar/service/jobqueue"

	"github.com/EricChiou/gooq"
)

var Event = eventRepository{}

type eventRepository struct{}

func (r eventRepository) AddEvent(e event.Entity) (sql.Result, error) {
	var result sql.Result
	err := jobqueue.Event.NewEventQueueJob(func() interface{} {
		sql := gooq.SQL{}
		sql.
			Insert(event.Table, event.Name, event.Type, event.StartTime, event.RepeatUnit, event.RepeatInterval, event.RepeatTime, event.LastTime, event.Remark, event.UserID).
			Values(e.Name, string(e.Type), e.StartTime, string(e.RepeatUnit), e.RepeatInterval, e.RepeatTime, e.LastTime, e.Remark, e.UserID)

		var err error
		result, err = db.Exec(sql.GetSQL())
		return err
	})

	return result, err
}

func (r eventRepository) GetEventListByUserID(userID int) (*sql.Rows, error) {
	var result *sql.Rows
	err := jobqueue.Event.NewEventQueueJob(func() interface{} {
		sql := gooq.SQL{}
		sql.Select("*").From(event.Table).Where(event.UserID).Eq(userID)

		var err error
		result, err = db.Query(sql.GetSQL())
		return err
	})

	return result, err
}
