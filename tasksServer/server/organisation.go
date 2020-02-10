package server

import (
	"../storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateOrganisation(w http.ResponseWriter, r *http.Request) {

	body, err, code := CheckGeneralRequest(http.MethodPost, w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(),code)
		return
	}
	createOrg := CreateOrganisationRequest{}

	err = json.Unmarshal(body, &createOrg)

	if err != nil {
		log.Println(string(body))
		log.Println(err)
		http.Error(w, "Ошибка преобразования запроса", 422)
		return
	}

	org := storage.Organisation{}

	org.Id = createOrg.Login + GetDateString()

	org.Name = createOrg.NameOrganisation


	org.IdAdmins = createOrg.Login + ";"
	org.IdUsers = createOrg.Login+";"

	err = storage.AddOrganisation(org)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Ошибка создания организации", 500)
		return
	}
	err = storage.AddOrganisationToUser(org.Id,createOrg.Login)
	if err != nil {
		fmt.Println(err)
		storage.DeleteOrganisation(org.Id)
		http.Error(w, "Ошибка создания организации", 500)
		return
	}
}
func GetOrganisations(w http.ResponseWriter, r * http.Request){

	body, err, code := CheckGeneralRequest(http.MethodPost, w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(),code)
		return
	}
	getOrg := GeneralRequest{}

	err = json.Unmarshal(body, &getOrg)
	if err != nil {
		log.Println(string(body))
		log.Println(err)
		http.Error(w, "Ошибка преобразования запроса", 422)
		return
	}
	user, err:= storage.GetUser(getOrg.Login)
	if err!=nil{
		http.Error(w, "Ошибка получения пользователя", 500)
		return
	}
	organisationsSlice:= StringToSlice(user.OrganisationsId.String)
	organisations:= GetOrganisationsResponse{}
	for _, v:= range organisationsSlice{
		org, _:= storage.GetOrganisation(v)
		organisations.Organisations = append(organisations.Organisations, org)
	}

	jsData, err:= json.Marshal(organisations)
	if err!=nil{
		http.Error(w, "Ошибка получения организаций", 422)
		return
	}
	w.Write(jsData)
}

