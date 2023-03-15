package entity

import "errors"

const (
	errEmptyEmail    = "err email is empty"
	errEmptyUserID   = "err user id is empty"
	errEmptyPassword = "err password is empty"
)

type Auth struct {
	ID       string
	UserID   string
	Password string
	Email    string
}

func (u *Auth) SetUserID(id string) error {
	u.UserID = id
	if id == "" {
		return errors.New(errEmptyUserID)
	}
	return nil
}

func (u *Auth) SetPassword(pass string) error {
	u.Password = pass
	if pass == "" {
		return errors.New(errEmptyPassword)
	}
	return nil
}

func (u *Auth) SetEmail(email string) error {
	u.Email = email
	if email == "" {
		return errors.New(errEmptyEmail)
	}
	return nil
}

func NewAuth(u *Auth) (*Auth, error) {
	//sanitize TODO

	return u, nil
}
