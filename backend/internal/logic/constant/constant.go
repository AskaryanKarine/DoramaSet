package constant

import (
	"fmt"
	"time"
)

const (
	EveryDayPoint    = 5
	EveryYearPoint   = 10
	LongNoLoginPoint = 50
	LongNoLoginHours = 4400.0
	LoginLen         = 5
	PasswordLen      = 8
	TokenExpiration  = time.Hour * 700
)

const (
	PublicList = iota
	PrivateList
)

var ListType = map[string]int{
	"public":  PublicList,
	"private": PrivateList,
}

func GetTypeList(val int) (string, error) {
	for k, v := range ListType {
		if v == val {
			return k, nil
		}
	}
	return "", fmt.Errorf("value doesn't exist in map")
}
