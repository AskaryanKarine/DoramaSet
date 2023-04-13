package db_erorrs

import "errors"

var (
	ErrorDontExistsInDB = errors.New("don't exists in db")
)
