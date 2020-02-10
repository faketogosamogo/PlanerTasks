package storage

import (

	"strings"
)

type Organisation struct{
	Id string
	IdUsers string
	IdDialogs string
	IdAdmins string
	Name string
	IdTasks string
}
func StringToSlice(str string)[]string{
	slice:= strings.Split(str,";")
	return slice
}
func SliceToString(slice []string)string{
	str:=""
	for _, value:= range slice{
		if value!="" {
			str += value + ";"
		}
	}
	return str
}
func AddOrganisation(org Organisation)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("insert into organisations values(?,?,?,?,?)",org.Id,org.IdUsers, org.IdDialogs, org.IdAdmins, org.Name,org.IdTasks)

	return err
}
func DeleteOrganisation(id string)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("delete from organisations where id=?",id)

	return err
}
func GetOrganisation(id string)(Organisation, error){
	db:= dbOpen()
	defer db.Close()
	org:= Organisation{}
	err:= db.QueryRow("select * from organisations where id=?", id).Scan(&org.Id,&org.IdUsers,&org.IdDialogs, &org.IdAdmins, &org.Name,&org.IdTasks)
	return org, err
}
func AddDialogToOrganisation(idOrganisation, idDialog string)error{

	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("update organisation.organisations set idDialogs = CONCAT(idDialogs,?) where id=?",idDialog+";", idOrganisation)

	return err
}
func DeleteDialogFromOrganisation(idOrganisation, idDialog string)error{
	db:=dbOpen()
	defer db.Close()
	org, err := GetOrganisation(idOrganisation)
	if err!=nil{
		return err
	}

	dialogs:= StringToSlice(org.IdDialogs)
	for i, v:= range dialogs{
		if v==idDialog{
			dialogs = append(dialogs[:i],dialogs[i+1])
			break
		}
	}
	dialogsString := SliceToString(dialogs)
	_, err = db.Exec("update organisation.organisations set idDialogs = ? where id=?",dialogsString, idOrganisation)
	return err

}