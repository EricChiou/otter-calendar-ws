package repository

import (
	"database/sql"
	"errors"
	"otter-calendar/app/model/user"
	"otter-calendar/app/types"
	"otter-calendar/db"
	"otter-calendar/service/jobqueue"

	"github.com/EricChiou/gooq"
)

var User = userRepository{}

type userRepository struct{}

func (r userRepository) AddUser(account, pwd, name string, role types.UserRole, status types.UserStatus) (sql.Result, error) {
	var result sql.Result
	err := jobqueue.User.NewUserQueueJob(func() interface{} {
		sql := gooq.SQL{}
		sql.
			Insert(user.Table, user.Account, user.Password, user.Name, user.Role, user.Status).
			Values(account, pwd, name, string(role), string(status))

		var err error
		result, err = db.Exec(sql.GetSQL())
		return err
	})

	return result, err
}

func (r userRepository) GetUser(account string) (user.Entity, error) {
	userEnt := user.Entity{}
	err := jobqueue.User.NewUserQueueJob(func() interface{} {
		sql := gooq.SQL{}
		sql.
			Select(user.ID, user.Account, user.Password, user.Name, user.Role, user.Status).
			From(user.Table).
			Where(user.Account).Eq(account)

		if row := db.QueryRow(sql.GetSQL()); row != nil {
			row.Scan(&userEnt.ID, &userEnt.Account, &userEnt.Password, &userEnt.Name, &userEnt.Role, &userEnt.Status)
		} else {
			return errors.New("no result")
		}

		return nil
	})

	return userEnt, err
}
