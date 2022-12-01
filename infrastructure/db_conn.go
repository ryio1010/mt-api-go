package infrastructure

import (
	"database/sql"
	"fmt"
	"mt-api-go/config"
)

const driverName = "postgres"

type PostgreSQLConnector struct {
	Conn *sql.DB
}

func NewPostgreSQLConnector() *PostgreSQLConnector {
	conf := config.LoadConfig()
	dsn := postgresConnInfo(*conf.PostgreSQLInfo)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &PostgreSQLConnector{
		Conn: conn,
	}
}

func postgresConnInfo(postgresInfo config.PostgreSQLInfo) string {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresInfo.User,
		postgresInfo.Password,
		postgresInfo.Host,
		postgresInfo.Port,
		postgresInfo.DbName,
	)
	fmt.Println(dataSourceName)

	return dataSourceName
}
