package storage

import "database/sql"

type Dialog struct {
	Id string
	UsersId string
	Name sql.NullString
	MessagesId sql.NullString
	IdCreator sql.NullString
	OrganisationId sql.NullString
}
func CreateDialog(d Dialog)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("insert into dialogs values(?,?,?,?,?,?)",d.Id,d.UsersId, d.Name, d.MessagesId, d.IdCreator, d.OrganisationId)

	return err
}

	func GetDialog(id string)(Dialog,error){
		db:=dbOpen()
		defer db.Close()

		d:=Dialog{}

		err:= db.QueryRow("select * from dialogs where id=?", id).Scan(&d.Id, &d.UsersId,&d.Name,&d.MessagesId,&d.IdCreator, &d.OrganisationId)

		return d,err
	}
func DeleteDialog(id string)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("delete from dialogs where id=?",id)

	return err
}
func AddMessageToDialog(idMessage, idDialog string) error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.dialogs set messagesid = CONCAT(messagesid,?) where id=?",idMessage+";", idDialog)

	return err
}
func AddUserToDialog(idUser, idDialog string) error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.dialogs set usersid = CONCAT(usersid,?) where id=?",idUser+";", idDialog)

	return err
}
