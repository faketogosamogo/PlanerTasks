package storage

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Task struct{
	Id string
	Description sql.NullString
	CreatorId string
	UsersId string
	DateStart mysql.NullTime
	DateFinish mysql.NullTime
	Status int
	IdDialog sql.NullString
	IdOrganisation sql.NullString
	Name sql.NullString
}