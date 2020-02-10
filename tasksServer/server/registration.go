package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../storage"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST", 405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", 400)
		return
	}
	req := RegistrationReguest{}
	err = json.Unmarshal(body, &req)

	if err != nil {
		http.Error(w, "Ошибка преобразования запроса", 422)
		return
	}
	user:= storage.User{}
	user.Login = req.Login
	user.Password = req.Password
	err = storage.AddUser(user)
	if err!=nil{
		fmt.Println(err)
		http.Error(w, "Ошибка добавления пользователя", 422)
		return
	}

	}