package clientHTTP

type GeneralRequest struct {
	Login string
}
type CreateOrganisationRequest struct {
	GeneralRequest
	NameOrganisation string
}
type GetTaskRequest struct{
	GeneralRequest
	IdTask string
}
type AuthorizationRequest struct {
	GeneralRequest
	Password string
}
type RegistrationReguest struct{
	GeneralRequest
	Password string
}
type UploadFileRequest struct{
	GeneralRequest
	FileExp string
	FileBody []byte
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