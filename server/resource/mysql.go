package resource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DatabaseMysql *sql.DB

func InitDatabase() (err error) {

	DatabaseMysql, err = sql.Open("mysql", "user:password@/database")
	return
}

func Close() {
	defer DatabaseMysql.Close()
}
