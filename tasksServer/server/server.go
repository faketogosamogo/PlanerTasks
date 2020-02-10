package server

import (
	"net/http"
)

var tokens map[string]string//[login]key
func StartServer(){
	/*auth:= AuthorizationRequest{}
	auth.Login = "g"
	auth.Password = "k"
	j,_:= json.Marshal(auth)*/
	//fmt.Println(string(j))
	//TO DO RAND
	tokens = make(map[string]string)
	http.HandleFunc("/dialog/create", CreateDialog)
	http.HandleFunc("/dialogs/get",GetDialogs)

	http.HandleFunc("/tasks/get",GetTasks)
	http.HandleFunc("/task/get",GetTask)
	http.HandleFunc("/task/create",CreateTask)
	http.HandleFunc("/task/addUser",AddUserToTask)
	http.HandleFunc("/task/deleteUser",DeleteUserFromTask)



	http.HandleFunc("/organisation/create", CreateOrganisation)
	http.HandleFunc("/organisations/get",GetOrganisations)

	http.HandleFunc("/auth",Authorization)
	http.HandleFunc("/registration",Registration)

	http.HandleFunc("/messages/send", SendMessage)
	http.HandleFunc("/messages/get",GetMessages)

	http.HandleFunc("/files/upload", UploadFile)
	http.HandleFunc("/files/get/",GetFiles)


	http.ListenAndServe("localhost:8080",nil)
}
