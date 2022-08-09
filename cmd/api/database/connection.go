package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	query := url.Values{}
	query.Add("database", conf.Database)
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(conf.User, conf.Password),
		Host:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		RawQuery: query.Encode(),
	}

	conn, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}
