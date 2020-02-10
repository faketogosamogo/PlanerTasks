package GUI

import (
	"../clientHTTP"
	"github.com/gotk3/gotk3/gtk"
)
var Token string
var Login string


func ShowAuthorizationWindow(){
	gtk.Init(nil)

	// Создаём билдер
	builder:= GetBuilderForFile("AuthorizationWindow.glade")

	obj:= GetObject("AuthorizationWindow", builder)
	win := obj.(*gtk.Window)

	obj = GetObject("btnAuthorization", builder)
	btnAuth:= obj.(*gtk.Button)

	obj = GetObject("tbLogin", builder)
	tbLogin:= obj.(*gtk.Entry)
	tbLogin.SetMaxLength(10)
	obj = GetObject("tbPassword", builder)
	tbPassword:= obj.(*gtk.Entry)
	tbPassword.SetMaxLength(10)
	obj= GetObject("btnRegistration", builder)
	btnRegistration:= obj.(*gtk.Button)

	btnAuth.Connect("clicked", func() {
		login,_:= tbLogin.GetText()
		password, _:= tbPassword.GetText()
		token, err:= clientHTTP.Authorization(login, password)
		if err!=nil{
			ShowMessage("Ошибка авторизации!", win)
			return
		}
		Token = token
		Login = login

		win.SetVisible(false)
		win.Destroy()

		ShowOrganisationWindow()


	})
	btnRegistration.Connect("clicked", func() {
		win.SetVisible(false)
		ShowRegistrationWindow()
		win.SetVisible(true)
	})
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}