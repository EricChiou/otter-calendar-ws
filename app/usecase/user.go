package usecase

import (
	"errors"
	"otter-calendar/app/repository"
	"otter-calendar/app/types"
	"otter-calendar/http/jwt"
	"otter-calendar/pkg/sha3"
)

var User = userUsecase{}

type userUsecase struct{}

func (u userUsecase) SignUp(account, pwd string) error {
	_, err := repository.User.AddUser(account, sha3.Encrypt(pwd, sha3.Long.L512), account, types.Normal, types.Active)
	return err
}

func (u userUsecase) Login(account, pwd string) (string, error) {
	userEnt, err := repository.User.GetUser(account)
	if err != nil {
		return "", err
	}

	if sha3.Encrypt(pwd, sha3.Long.L512) != userEnt.Password {
		return "", errors.New("password error")
	}

	token := jwt.Generate(userEnt.ID, userEnt.Account, userEnt.Name, userEnt.Role, userEnt.Status)
	return token, err
}
