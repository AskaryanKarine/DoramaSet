package errors

import (
	"errors"
	"fmt"
)

type BalanceError struct {
	Have int
	Want int
}

func (b BalanceError) Error() string {
	return fmt.Sprintf("insufficient funds: you have %d, except %d", b.Have, b.Want)
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
)
