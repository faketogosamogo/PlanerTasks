package GUI

import (
	"../clientHTTP"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
)
func ShowTaskWindow(){

	gtk.Init(nil)
	builder:= GetBuilderForFile("TaskWindow.glade")

	obj:=GetObject("TaskWindow",builder)
	win:= obj.(*gtk.Window)

	task, err:= clientHTTP.GetTask(Login,idTask,Token)
	if err!=nil{
		ShowMessage("Ошибка получения задачи", win)
		win.Destroy()
	}

	obj =GetObject("tbId",builder)
	tbId:= obj.(*gtk.Entry)
	tbId.SetText(task.Id)

	obj =GetObject("tbDescription",builder)
	tbDescription:= obj.(*gtk.TextView)
	textView,_:= gtk.TextViewNew()
	buffer, _:= textView.GetBuffer()
	buffer.SetText(task.Description.String)
	tbDescription.SetBuffer(buffer)

	obj =GetObject("tbCreatorId",builder)
	tbCreatorId:= obj.(*gtk.Entry)
	tbCreatorId.SetText(task.CreatorId)

	obj =GetObject("tbDateStart",builder)
	tbDateStart:= obj.(*gtk.Entry)
	tbDateStart.SetText(task.DateStart.Time.String())

	obj =GetObject("tbDateFinish",builder)
	tbDateFinish:= obj.(*gtk.Entry)
	tbDateFinish.SetText(task.DateFinish.Time.String())

	obj =GetObject("status",builder)
	status:= obj.(*gtk.SpinButton)
	status.SetText(strconv.Itoa(task.Status))


	obj = GetObject("tbIdDialog",builder)
	tbIdDialog:= obj.(*gtk.Entry)
	tbIdDialog.SetText(task.IdDialog.String)

	obj = GetObject("tbIdOrganisation",builder)
	tbIdOrganisation:= obj.(*gtk.Entry)
	tbIdOrganisation.SetText(task.IdOrganisation.String)

	obj = GetObject("tbName",builder)
	tbName:= obj.(*gtk.Entry)
	tbName.SetText(task.Name.String)

	obj = GetObject("btnDeleteUser",builder)
	btnDeleteUser:= obj.(*gtk.Button)

	obj = GetObject("btnAddUser",builder)
	btnAddUser:= obj.(*gtk.Button)

	obj = GetObject("btnSaveStatus",builder)
	btnSaveStatus:= obj.(*gtk.Button)

	obj = GetObject("tbLoginUser",builder)
	tbLoginUser:= obj.(*gtk.Entry)

	obj = GetObject("btnDeleteTask",builder)
	btnDeleteTask:= obj.(*gtk.Button)

	if Login!=task.CreatorId{
		btnAddUser.SetVisible(false)
		btnDeleteUser.SetVisible(false)
		tbLoginUser.SetVisible(false)
		btnSaveStatus.SetVisible(false)
	}

	obj =GetObject("btnSaveStatus",builder)
	//btnSaveStatus:= obj.(*gtk.Button)

	if Login==task.CreatorId{
		status.SetEditable(true)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
		ShowMainWindow()
	})
	win.ShowAll()
	if Login!=task.CreatorId{
		btnAddUser.SetVisible(false)
		btnDeleteUser.SetVisible(false)
		tbLoginUser.SetVisible(false)
		btnSaveStatus.SetVisible(false)
		btnDeleteTask.SetVisible(false)
	}
	gtk.Main()


}
