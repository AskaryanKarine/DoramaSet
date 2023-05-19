package errors

import (
	"errors"
	"fmt"
)

type BalanceError struct {
	Have int `json:"have"`
	Want int `json:"want"`
}

func (b BalanceError) Error() string {
	return fmt.Sprintf("you have %d scores, but you need %d scores", b.Have, b.Want)
}

type LoginLenError struct {
	LoginLen int
}

func (l LoginLenError) Error() string {
	return fmt.Sprintf("login must be more then %d symbols", l.LoginLen)
}

type PasswordLenError struct {
	PasswordLen int
}

func (p PasswordLenError) Error() string {
	return fmt.Sprintf("password must be more then %d symbols", p.PasswordLen)
}

var (
	ErrorAdminAccess    = errors.New("low level of access")
	ErrorCreatorAccess  = errors.New("no access right")
	ErrorInvalidEmail   = errors.New("invalid email")
	ErrorWrongLogin     = errors.New("wrong login or password")
	ErrorUserExist      = errors.New("user already exists")
	ErrorDontExistsInDB = errors.New("don't exists in db")
	ErrorPublic         = errors.New("doesn't public list")
)
