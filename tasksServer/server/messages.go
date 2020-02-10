package server

import (
	"../storage"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SendMessage(w http.ResponseWriter, r * http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	sendMessage := SendMessageRequest{}
	err = json.Unmarshal(body, &sendMessage)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}
	dialog, err:= storage.GetDialog(sendMessage.IdDialog)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка обращения к диалогу", 422)
		return
	}

	users:= StringToSlice(dialog.UsersId)
	if !CheckUserInDialog(users, sendMessage.Login){
		http.Error(w, "Ошибка доступа к диалогу", 403)
		return
	}

	message:= storage.Message{}
	message.Id = sendMessage.Login+GetDateString()+"m"
	message.DialogId = sendMessage.IdDialog
	message.Time.Time = time.Now()
	message.Time.Valid = true
	message.Type = sendMessage.Type
	message.Text = sendMessage.Text
	message.LoginSender = sendMessage.Login


	err = storage.AddMessage(message)
	if err!=nil{
		log.Println(err.Error())
		http.Error(w, "Ошибка создания сообщения", 500)
		return
	}
	err = storage.AddMessageToDialog(message.Id, sendMessage.IdDialog)
	if err!=nil{
		storage.DeleteMessage(message.Id)
		http.Error(w, "Ошибка добавления сообщения", 500)
		return
	}
	}
func GetMessages(w http.ResponseWriter, r *http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	getMessages:= GetMessagesRequest{}
	err = json.Unmarshal(body, &getMessages)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}

	dialog, err:= storage.GetDialog(getMessages.IdDialog)
	if err!=nil{
		http.Error(w, "Ошибка получения организации", 422)
		return
	}
	messagesSlice := StringToSlice(dialog.MessagesId.String)

	messagesResp :=GetMessagesResponse{}

	for _, value:= range messagesSlice{
		message, _:= storage.GetMessage(value)
		messagesResp.Messages = append(messagesResp.Messages, message)
	}
	jsData, err:= json.Marshal(messagesResp)
	if err!=nil{
		http.Error(w, "Ошибка получения организаций", 422)
		return
	}

	w.Write(jsData)
}


