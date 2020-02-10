package storage

import (
	"database/sql"

)


type Dialog struct {
	Id string
	UsersId string
	Name sql.NullString
	MessagesId sql.NullString
	IdCreator sql.NullString
	OrganisationId sql.NullString
}
