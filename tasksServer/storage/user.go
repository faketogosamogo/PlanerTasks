package storage

import "database/sql"

type User struct {
	Login string
	Password string
	DialogsId sql.NullString
	OrganisationsId sql.NullString
	Tasks sql.NullString
}
func AddUser(user User)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("insert into users values(?,?,?,?,?)",user.Login, user.Password, user.DialogsId, user.OrganisationsId, user.Tasks)
	return err
}
func DeleteUser(login string)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("delete from organisation.users where login=?",login)
	return err
}
func GetUser(login string)(User,error){
	db:=dbOpen()
	defer db.Close()

	u:=User{}

	err:= db.QueryRow("select * from users where login=?", login).Scan(&u.Login, &u.Password, &u.DialogsId, &u.OrganisationsId,&u.Tasks)

	return u,err

}
func AddTaskToUser(idTask, idUser string)error{

	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.users set tasksId = CONCAT_WS('',tasksId,?) where login=?",idTask+";", idUser)

	return err
}
func DeleteTaskFromUser(idTask, idUser string)error{
	db:=dbOpen()
	defer db.Close()

	user, err := GetUser(idUser)
	if err!=nil{
		return err
	}

	tasks:= StringToSlice(user.Tasks.String)
	for i, v:= range tasks{
		if v==idTask{
			tasks = append(tasks[:i],tasks[i+1])
			break
		}
	}
	tasksSlice := SliceToString(tasks)
	_, err = db.Exec("update organisation.users set tasksId = ? where login = ?",tasksSlice, idUser)
	return err

}
func AddDialogToUser(idDialog, idUser string)error{

	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.users set dialogsId = CONCAT_WS('',dialogsId,?) where login=?",idDialog+";", idUser)

	return err
}
func DeleteDialogFromUser(idDialog, idUser string)error{
	db:=dbOpen()
	defer db.Close()

	user, err := GetUser(idUser)
	if err!=nil{
		return err
	}

	dialogs:= StringToSlice(user.DialogsId.String)
	for i, v:= range dialogs{
		if v==idDialog{
			dialogs = append(dialogs[:i],dialogs[i+1])
			break
		}
	}
	dialogsSlice := SliceToString(dialogs)
	_, err = db.Exec("update organisation.users set dialogsId = ? where login = ?",dialogsSlice, idUser)
	return err
}
func AddOrganisationToUser(idOrganisation, idUser string)error{

	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.users set organisationsId =  CONCAT_WS('',organisationsId,?) where login=?",idOrganisation+";", idUser)

	return err
}
func DeleteOrganisationFromUser(idDialog, idUser string)error{
	db:=dbOpen()
	defer db.Close()

	user, err := GetUser(idUser)
	if err!=nil{
		return err
	}

	organisations:= StringToSlice(user.OrganisationsId.String)
	for i, v:= range organisations{
		if v==idDialog{
			organisations = append(organisations[:i],organisations[i+1])
			break
		}
	}
	organisationSlice := SliceToString(organisations)
	_, err = db.Exec("update organisation.users set dialogsId = ? where login = ?",organisationSlice, idUser)
	return err
}