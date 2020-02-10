package clientHTTP

import (
	"../storage"
	"time"
)
const(
	URL = "http://127.0.0.1:8080"

	AuthURL = URL+"/auth"
	RegistrationURL= URL+ "/registration"

	GetOrganisationsURL = URL+"/organisations/get"

	GetDialogsURL = URL+"/dialogs/get"
	CreateDialogURL = URL + "/dialog/create"

	GetTasksURL = URL+"/tasks/get"
	GetTaskURL = URL + "/task/get"
	AddUserToTaskURL = URL+"/task/addUser"
	DeleteUserFromTaskURL = URL+"/task/deleteUser"

	SendMessageURL = URL+"/messages/send"
	GetMessagesURL = URL+"/messages/get"

	FileGetURL = URL+"/files/get/"
	FileUploadURL = URL + "/files/upload"
	)


type GeneralResponse struct {
	Time time.Time

}
type AuthorizationResponse struct {
	GeneralResponse
	Token string
}
type GetOrganisationsResponse struct{
	Organisations []storage.Organisation
}
type GetDialogsResponse struct{
	Dialogs []storage.Dialog
}
type GetMessagesResponse struct{
	Messages []storage.Message
}
type GetTasksResponse struct {
	Tasks []storage.Task
}