package db

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	fmt.Println(config.StringConnection)
	conn, err := sql.Open("mysql", config.StringConnection)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		fmt.Println(err)
		conn.Close()
		return nil, err
	}

	return conn, nil
}
