package config

type dbConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	Port     int
	Type     string
}

type appConfig struct {
	SecretKey            string
	EveryDayPoint        int
	EveryYearPoint       int
	LongNoLoginPoint     int
	LongNoLoginHours     float64
	LoginLen             int
	PasswordLen          int
	TokenExpirationHours int
}

type loggerConfig struct {
	FileName string
	Level    string
}

type serverConfig struct {
	Port string
	Mode string
}

type OTConfig struct {
	Endpoint    string
	ServiceName string
	Ratio       float64
}

type Config struct {
	DB            dbConfig
	App           appConfig
	Logger        loggerConfig
	Server        serverConfig
	OpenTelemetry OTConfig
}
