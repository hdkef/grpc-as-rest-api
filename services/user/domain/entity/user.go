package entity

import (
	"errors"
)

const (
	errEmptyID       = "err user id is empty"
	errEmptyEmail    = "err email is empty"
	errEmptyName     = "err name is empty"
	errEmptyAddress  = "err address is empty"
	errEmptyPassword = "err password is empty"
)

type User struct {
	ID       string
	Email    string
	Name     string
	Address  string
	Password string
}

func (u *User) SetID(id string) error {
	if id == "" {
		return errors.New(errEmptyID)
	}
	u.ID = id
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New(errEmptyEmail)
	}
	u.Email = email
	return nil
}

func (u *User) SetAddress(address string) error {
	if address == "" {
		return errors.New(errEmptyAddress)
	}
	u.Address = address
	return nil
}

func (u *User) SetName(name string) error {
	if name == "" {
		return errors.New(errEmptyName)
	}
	u.Name = name
	return nil
}

func (u *User) SetPassword(pass string) error {
	if pass == "" {
		return errors.New(errEmptyPassword)
	}

	u.Password = pass
	return nil
}

func NewUser(u *User) (*User, error) {
	//sanitize TODO

	return u, nil
}
