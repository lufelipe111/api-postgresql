package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/study/api-postgresql/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err) // don't use on production environment
	}

	err = conn.Ping() // Ping to verify if the connection was established

	return conn, err
}
