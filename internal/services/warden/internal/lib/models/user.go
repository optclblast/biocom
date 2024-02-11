package models

import "time"

type UserIdentity interface {
	ID() string
	OrganizationID() string
	IsDeleted() bool
	IsAdmin() bool
}

type User struct {
	Id             string
	OrganizationId string
	Login          string
	Name           string
	PasswordHash   []byte
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Deleted        bool
	Admin          bool
}

func (u *User) ID() string {
	return u.Id
}

func (u *User) OrganizationID() string {
	return u.OrganizationId
}

func (u *User) IsDeleted() bool {
	return u.Deleted
}

func (u *User) IsAdmin() bool {
	return u.Admin
}
