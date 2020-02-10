package clientHTTP

import (
	"../storage"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDialogs(login,idorganisation, token string) ([]storage.Dialog, error){
	getDialogs:= GetDialogsRequest{}
	getDialogs.Login = login
	getDialogs.IdOrganisation = idorganisation

	js, err := json.Marshal(getDialogs)
	if err!=nil{
		return nil, err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)


	req, err:= http.NewRequest("POST", GetDialogsURL,&buf)
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
	dialogs:= GetDialogsResponse{}
	err = json.Unmarshal(body, &dialogs)
	return dialogs.Dialogs, err
}
func CreateDialog(login, idorganisation,nameDialog, token string) error{
	createDialog:= CreateDialogRequest{}
	createDialog.Login = login
	createDialog.IdOrganisation = idorganisation
	createDialog.NameDialog = nameDialog
	js, err := json.Marshal(createDialog)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)


	req, err:= http.NewRequest("POST", CreateDialogURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode!=200{
		return errors.New("Ошибка создания диалога!")
	}
	return nil
}