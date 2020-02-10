package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



const (
	driverName = "mysql"
	connectionString = "root:1234@/organisation"
)
func dbOpen() (*sql.DB){
	db, err:=sql.Open(driverName, connectionString)
	if err!=nil{
		panic(err)
	}

	return db
}