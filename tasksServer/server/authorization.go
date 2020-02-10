package server

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"../storage"
	"time"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lenToken = 8
func GenerationToken(n int)string{
	str:=make([]rune, n)
	for i,_:=range str{
		str[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(str)
}

func Authorization(w http.ResponseWriter, r *http.Request){
	if r.Method!= http.MethodPost{
		http.Error(w,"POST",405)
		return
	}
	body, err:= ioutil.ReadAll(r.Body)

	if err!=nil{
		http.Error(w, "Ошибка чтения тела запроса", 400)
		return
	}
	req := AuthorizationRequest{}
	err = json.Unmarshal(body, &req)

	if err!=nil{
		http.Error(w,"Ошибка преобразования запроса",422)
		return
	}
	user, err:= storage.GetUser(req.Login)
	if err!=nil{
		log.Println(err)
		if err==sql.ErrNoRows{
			http.Error(w,"Неверные данные для входа",400)
			return
		}
		http.Error(w,"Ошибка сервера",500)
		return
	}
	if user.Login!=req.Login || user.Password!=req.Password{
		http.Error(w,"Неверные данные для входа",400)
		return
	}
	token:= GenerationToken(lenToken)
	tokens[user.Login] = token

	authResp:=AuthorizationResponse{}
	authResp.Time =  time.Now()

	authResp.Token = token

	resp,_:= json.Marshal(authResp)
	w.Write(resp)
}
