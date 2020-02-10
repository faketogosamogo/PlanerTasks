package server

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type GeneralRequest struct {
	Login string
}
type CreateOrganisationRequest struct {
	GeneralRequest
	NameOrganisation string
}
type UploadFileRequest struct{
	GeneralRequest
	FileExp string
	FileBody []byte
}

type AuthorizationRequest struct {
	GeneralRequest
	Password string
}
type CreateDialogRequest struct {
	GeneralRequest
	NameDialog string
	IdOrganisation string
	Users []string
}
type SendMessageRequest struct {
	GeneralRequest
	IdDialog string
	Text string
	Type string
}
type GetDialogsRequest struct{
	GeneralRequest
	IdOrganisation string
}
type GetMessagesRequest struct{
	GeneralRequest
	IdDialog string
}
type GetTasksRequest struct {
	GeneralRequest
	IdOrganisation string
}
type EditTaskRequest struct{
	GeneralRequest
	TaskId string
	UserLogin string
}
type GetTaskRequest struct{
	GeneralRequest
	IdTask string
}
type RegistrationReguest struct{
	GeneralRequest
	Password string
}
type CreateTasksRequest struct{
	GeneralRequest
	Description sql.NullString
	DateFinish mysql.NullTime
	IdOrganisation string
	Name sql.NullString
}
