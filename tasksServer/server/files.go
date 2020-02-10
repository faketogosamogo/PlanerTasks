package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

)
func UploadFile(w http.ResponseWriter, r *http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		log.Println(err)
		http.Error(w, err.Error(),code)
		return
	}
	uploadFile:= UploadFileRequest{}
	err = json.Unmarshal(body, &uploadFile)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}
	filename:= uploadFile.Login + GetDateString() + "."+uploadFile.FileExp

	file, err:= os.Create("./files/"+filename)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка создания файла", 500)
		return
	}
	_, err = file.Write(uploadFile.FileBody)
	if err!=nil{
		log.Println(err)
		os.Remove(filename)
		http.Error(w, "Ошибка записи файла", 422)
		return
	}
	w.Write([]byte(filename))
}
func GetFiles(w http.ResponseWriter, r *http.Request){
	if r.Method!= http.MethodGet{
		http.Error(w,"GET",405)
		return
	}
	path:= r.URL.Path
	pathSlice:= strings.Split(path,"/")
	file, err:= os.Open("./files/"+pathSlice[len(pathSlice)-1])

	if err!=nil{
		fmt.Println(err)
		http.Error(w, "Файл не найден", 422)
		return
	}
	fileBody, err:= ioutil.ReadAll(file)
	if err!=nil{
		http.Error(w, "Ошибка чтения файла", 500)
		return
	}
	w.Write(fileBody)
}
