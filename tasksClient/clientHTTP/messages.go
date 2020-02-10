package clientHTTP

import (
	"../storage"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetMessages(login,idDialog, token string) ([]storage.Message, error){
	getMessages := GetMessagesRequest{}
	getMessages.IdDialog = idDialog
	getMessages.Login = login

	js, err := json.Marshal(getMessages)
	if err!=nil{
		return nil, err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)


	req, err:= http.NewRequest("POST", GetMessagesURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil{
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, err
	}
	messages:= GetMessagesResponse{}
	err = json.Unmarshal(body, &messages)
	return messages.Messages, err
}

func SendMessage(login, idDialog, token , text, typeMessage string)error{
	sendMess:= SendMessageRequest{}
	sendMess.Login = login
	sendMess.Text = text
	sendMess.IdDialog = idDialog
	sendMess.Type = typeMessage
	js, err := json.Marshal(sendMess)
	if err!=nil{
		return err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)

	req, err:= http.NewRequest("POST", SendMessageURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil{
		return  err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode!=200{
		return errors.New("Ошибка отправки сообщения!")
	}
	return nil
}