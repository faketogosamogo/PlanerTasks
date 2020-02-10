package clientHTTP

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func Registration(login, password string)(error){
	reg:= RegistrationReguest{GeneralRequest{login},password}

	jsReg, _:= json.Marshal(reg)
	buf:= bytes.Buffer{}

	buf.Write(jsReg)

	resp, err:= http.Post(RegistrationURL,"application/json",&buf )
	if err!=nil{
		return err
	}

	if resp.StatusCode!=200{
		return errors.New("Ошибка создания!")
	}

	return nil
	}
