package server

import (
	"../storage"
	"time"
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