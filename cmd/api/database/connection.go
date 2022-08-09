package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func OpenConnection() (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", "CDC_POC")
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "@docker@2021"),
		Host:     fmt.Sprintf("%s:%s", "localhost", "1433"),
		RawQuery: query.Encode(),
	}

	conn, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}
