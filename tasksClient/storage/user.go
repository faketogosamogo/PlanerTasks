package storage

import "database/sql"

type User struct {
	Login string
	Password string
	DialogsId sql.NullString
	OrganisationsId sql.NullString
	Tasks sql.NullString

}
