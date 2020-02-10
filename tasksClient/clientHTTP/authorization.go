package clientHTTP

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Authorization(login, password string)(string, error){
	auth:= AuthorizationRequest{}
	auth.Login = login
	auth.Password = password

	jsAuth, _:= json.Marshal(auth)
	buf:= bytes.Buffer{}

	buf.Write(jsAuth)

	resp, err:= http.Post(AuthURL,"application/json",&buf )
	if err!=nil{
		return "",err
	}
	body, err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		return "", err
	}
	authResp:=AuthorizationResponse{}
	err = json.Unmarshal(body, &authResp)
	if err!=nil{
		return "", err
	}

	return authResp.Token, nil

	}
