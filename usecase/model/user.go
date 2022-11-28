package model

import "mt-api-go/domain/model"

type userId string

type User struct {
	ID       userId `json:"userid"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func UserFromDomainModel(m *model.MUser) *User {
	u := &User{
		ID:       userId(m.Userid),
		Name:     m.Username,
		Password: m.Password,
	}

	return u
}
