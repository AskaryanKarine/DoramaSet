package config

type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	Port     int
}

type App struct {
	SecretKey            string
	EveryDayPoint        int
	EveryYearPoint       int
	LongNoLoginPoint     int
	LongNoLoginHours     float64
	LoginLen             int
	PasswordLen          int
	TokenExpirationHours int
}

type Config struct {
	DB  DB
	App App
}
