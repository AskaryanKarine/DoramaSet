package constant

import "time"

const (
	EveryDayPoint    = 5
	EveryYearPoint   = 10
	LongNoLoginPoint = 50
	LongNoLoginHours = 4400.0
	LoginLen         = 5
	PasswordLen      = 8
	TokenExpiration  = time.Hour * 700
	PublicType       = "public"
)
