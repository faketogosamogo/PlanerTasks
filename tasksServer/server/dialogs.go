package server

import (
	"../storage"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
func CheckUserInDialog(users []string, user string)bool{
	for _, value := range users{
		if value==user{
			return true
		}
	}
	return false
}
func CreateDialog(w http.ResponseWriter, r *http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	createDialog:= CreateDialogRequest{}
	err = json.Unmarshal(body, &createDialog)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}


	organisation, err:= storage.GetOrganisation(createDialog.IdOrganisation)
	if err!=nil{
		fmt.Println(err)
		http.Error(w, "Ошибка обращения к организации", 422)
		return
	}
	users:= StringToSlice(organisation.IdUsers)
	if !CheckUserInDialog(users, createDialog.Login){
		http.Error(w, "Вы не имеете доступа к данной организации!", 403)
		return
	}


	dialog := storage.Dialog{}
	dialog.Id = createDialog.Login + GetDateString()
	dialog.OrganisationId = sql.NullString{createDialog.IdOrganisation,true}
	dialog.IdCreator.String = createDialog.Login
	dialog.UsersId +=dialog.IdCreator.String+";"
	dialog.UsersId += SliceToString(createDialog.Users)
	dialog.Name.String = createDialog.NameDialog
	dialog.Name.Valid = true


	err = storage.CreateDialog(dialog)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка создания диалого", 500)
		return
	}

	err = storage.AddDialogToOrganisation(createDialog.IdOrganisation, dialog.Id)
	if err!=nil{
		log.Println(err)
		storage.DeleteDialog(dialog.Id)
		http.Error(w, "Ошибка добавления диалога", 500)
		return
	}
	err = storage.AddDialogToUser(dialog.Id, createDialog.Login)
	if err!=nil{
		log.Println(err)
		storage.DeleteDialog(dialog.Id)
		storage.DeleteDialogFromOrganisation(dialog.Id, organisation.Id)
		http.Error(w, "Ошибка добавления диалога", 500)
		return
	}
	}

func GetDialogs(w http.ResponseWriter, r *http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	getDialogs:= GetDialogsRequest{}
	err = json.Unmarshal(body, &getDialogs)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}

	user, err:= storage.GetUser(getDialogs.Login)
	if err!=nil{
		http.Error(w, "Ошибка получения организации", 422)
		return
	}
	dialogsSlice := StringToSlice(user.DialogsId.String)

	dialogsResp :=GetDialogsResponse{}

	for _, value:= range dialogsSlice{
		dialog, _:= storage.GetDialog(value)

		if dialog.OrganisationId.String==getDialogs.IdOrganisation{
			dialogsResp.Dialogs = append(dialogsResp.Dialogs, dialog)
		}
	}
	jsData, err:= json.Marshal(dialogsResp)
	if err!=nil{
		http.Error(w, "Ошибка получения организаций", 422)
		return
	}

	w.Write(jsData)
}