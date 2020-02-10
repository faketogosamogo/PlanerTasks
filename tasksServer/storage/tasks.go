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
func AddUserToTask(idUser, idTask string) error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.tasks set usersid = CONCAT(usersid,?) where id=?",idUser+";", idTask)

	return err
}
func AddTask(t Task)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("insert into tasks values(?,?,?,?,?,?,?,?,?)",t.Id,t.Description,t.CreatorId,t.UsersId,t.DateStart,t.DateFinish,t.Status,t.IdDialog, t.IdOrganisation, t.Name)

	return err
}
func GetTask(id string)(Task,error){
	db:=dbOpen()
	defer db.Close()
	t:= Task{}
	err:= db.QueryRow("select * from tasks where id=?", id).Scan(&t.Id,&t.Description,&t.CreatorId,&t.UsersId,&t.DateStart,&t.DateFinish,&t.Status,&t.IdDialog, &t.IdOrganisation, &t.Name)
	return t, err
}
func DeleteTask(id string)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("delete from tasks where id=?",id)

	return err
}
func DeleteUserFromTask(idUser, idTask string)error{
	db:=dbOpen()
	defer db.Close()
	task, err := GetTask(idTask)
	if err!=nil{
		return err
	}

	users:= StringToSlice(task.UsersId)
	for i, v:= range users{
		if v==idUser{
			users = append(users[:i],users[i+1])
			break
		}
	}
	usersString := SliceToString(users)
	_, err = db.Exec("update organisation.tasks set usersid = ? where id=?",usersString, idTask)
	return err

}