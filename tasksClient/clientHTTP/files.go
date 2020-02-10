package clientHTTP

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(login, path, token string)(string, error){
	uplFile:= UploadFileRequest{}
	uplFile.Login = login

	file, err:= os.Open(path)
	if err!=nil{
		return "", err
	}
	exp:=filepath.Ext(path)
	fmt.Println(exp)
	uplFile.FileExp = exp

	fileBody, err:= ioutil.ReadAll(file)
	if err!=nil{
		return "", err
	}
	uplFile.FileBody = fileBody

	js, err := json.Marshal(uplFile)
	if err!=nil{
		return "" ,err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)

	req, err:= http.NewRequest("POST", FileUploadURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil{
		return "",  err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode!=200{
		return "", errors.New("Ошибка отправки")
	}
	body, err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		return "", errors.New("Ошибка чтения ответа")
	}
	return string(body), nil
}