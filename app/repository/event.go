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
			Insert(event.Table, event.Name, event.Type, event.StartTime, event.RepeatUnit, event.RepeatInterval, event.RepeatTime, event.LastTime, event.Remark, event.UserID, event.CalType).
			Values(e.Name, string(e.Type), e.StartTime, string(e.RepeatUnit), e.RepeatInterval, e.RepeatTime, e.LastTime, e.Remark, e.UserID, string(e.CalType))

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

func (r eventRepository) GetEventByEventID(eventID, userID int) *sql.Row {
	var result *sql.Row
	jobqueue.Event.NewEventQueueJob(func() interface{} {
		sql := gooq.SQL{}
		sql.Select("*").From(event.Table).Where(event.ID).Eq(eventID).And(event.UserID).Eq(userID)

		result = db.QueryRow(sql.GetSQL())
		return nil
	})

	return result
}

func (r eventRepository) UpdateEvent(e event.Entity) (sql.Result, error) {
	var result sql.Result
	err := jobqueue.Event.NewEventQueueJob(func() interface{} {
		sql := gooq.SQL{}
		conditions := []string{
			gooq.Column(event.Name).Eq(e.Name),
			gooq.Column(event.Type).Eq(string(e.Type)),
			gooq.Column(event.StartTime).Eq(e.StartTime),
			gooq.Column(event.RepeatUnit).Eq(string(e.RepeatUnit)),
			gooq.Column(event.RepeatInterval).Eq(e.RepeatInterval),
			gooq.Column(event.RepeatTime).Eq(e.RepeatTime),
			gooq.Column(event.LastTime).Eq(e.LastTime),
			gooq.Column(event.Remark).Eq(e.Remark),
			gooq.Column(event.CalType).Eq(e.CalType),
		}
		sql.
			Update(event.Table).
			Set(conditions...).
			Where(event.ID).Eq(e.ID).And(event.UserID).Eq(e.UserID)

		var err error
		result, err = db.Exec(sql.GetSQL())
		return err
	})

	return result, err
}

func (r eventRepository) DeleteEventByEventID(eventID, userID int) (sql.Result, error) {
	var result sql.Result
	err := jobqueue.Event.NewEventQueueJob(func() interface{} {
		sql := gooq.SQL{}
		sql.Delete(event.Table).Where(event.ID).Eq(eventID).And(event.UserID).Eq(userID)

		var err error
		result, err = db.Exec(sql.GetSQL())
		return err
	})

	return result, err
}
