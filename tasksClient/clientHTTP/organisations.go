package clientHTTP

import (
	"../storage"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
func GetOrganisations(login, token string) ([]storage.Organisation, error){
	getOrg:= GeneralRequest{}
	getOrg.Login = login

	js, _ := json.Marshal(getOrg)
	buf:= bytes.Buffer{}
	buf.Write(js)


	req, err:= http.NewRequest("POST", GetOrganisationsURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil{
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, err
	}
	organisations:= GetOrganisationsResponse{}
	err = json.Unmarshal(body, &organisations)
	return organisations.Organisations, err
}
