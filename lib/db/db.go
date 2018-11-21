package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	database = "test"
	user     = "root"
	password = "1234"
)

var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

//InitDB 建立一個共用的DB物件給大家用
func InitDB() *sql.DB {
	costdb, _ := sql.Open("mysql", connectionString)
	return costdb
}
