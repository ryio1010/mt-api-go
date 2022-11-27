package config

const (
	port       = "8081"
	dbUser     = "mtadmin"
	dbPassword = "mtadmin"
	dbName     = "muscletracking"
	dbHost     = "localhost"
	dbPort     = "5433"
	dbSchema   = "muscle_tracking_go"
)

type AppConfig struct {
	HTTPInfo       *HttpInfo
	PostgreSQLInfo *PostgreSQLInfo
}

type HttpInfo struct {
	Addr string
}

type PostgreSQLInfo struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     string
	Schema   string
}

func LoadConfig() *AppConfig {
	addr := ":" + port

	httpInfo := &HttpInfo{
		Addr: addr,
	}

	dbInfo := &PostgreSQLInfo{
		User:     dbUser,
		Password: dbPassword,
		DbName:   dbName,
		Host:     dbHost,
		Port:     dbPort,
		Schema:   dbSchema,
	}

	conf := AppConfig{
		HTTPInfo:       httpInfo,
		PostgreSQLInfo: dbInfo,
	}

	return &conf
}
