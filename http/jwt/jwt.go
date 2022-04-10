package jwt

import (
	"otter-calendar/config"
	"otter-calendar/pkg/jwt"
	"time"
)

var alg = jwt.Alg.HS256

type Payload struct {
	ID     int    `json:"id"`
	Acc    string `json:"acc"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status string `json:"status"`
	Exp    int64  `json:"exp"`
}

func Generate(id int, acc, name, role, status string) string {
	payload := Payload{
		ID:     id,
		Acc:    acc,
		Name:   name,
		Role:   role,
		Status: status,
		Exp:    (time.Now().Unix() + int64(config.Get().JWTExpire*86400)) * 1000,
	}

	token, _ := jwt.Generate(payload, config.Get().JWTKey, alg)
	return token
}

func Verify(j, k string) (Payload, error) {
	var payload Payload
	err := jwt.Verify(&payload, j, config.Get().JWTKey, alg)

	return payload, err
}
