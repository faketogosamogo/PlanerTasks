package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"net/http"
)

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
func GetDateString()string{
	now:= time.Now()

	return strconv.Itoa(now.Year())+strconv.Itoa(int(now.Month()))+ strconv.Itoa(now.Day())+strconv.Itoa(now.Hour())+strconv.Itoa(now.Minute())+strconv.Itoa(now.Second())
}
func CheckGeneralRequest(method string, w http.ResponseWriter, r* http.Request)([]byte,error,int){

	if r.Method!= method{
		http.Error(w,method,405)
		return nil,errors.New("Не подходящий метод"),405
	}
	body, err:= ioutil.ReadAll(r.Body)

	if err!=nil{

		http.Error(w, "Ошибка чтения тела запроса", 400)
		return nil,errors.New("Ошибка чтения тела запроса"),400
	}

	req:=GeneralRequest{}
	err = json.Unmarshal(body, &req)

	if err!=nil{

		http.Error(w,"Ошибка преобразования запроса",422)
		return nil,errors.New("Ошибка преобразования запроса"),422
	}

	token:= r.Header.Get("Access-Token")

	if token!=tokens[req.Login]{
		http.Error(w, "Ошибка авторизации",401)
		return nil,errors.New("Ошибка авторизации"),401
	}

	return body,nil,0
}
