package GUI

import (
	"../clientHTTP"
	"github.com/gotk3/gotk3/gtk"
)

func ShowRegistrationWindow(){
	gtk.Init(nil)
	builder:= GetBuilderForFile("RegistrationWindow.glade")

	obj:=GetObject("RegistrationWindow",builder)
	win:= obj.(*gtk.Window)

	obj = GetObject("entryLogin", builder)
	tbLogin := obj.(*gtk.Entry)
	tbLogin.SetMaxLength(10)

	obj = GetObject("entryPassword", builder)
	tbPassword := obj.(*gtk.Entry)
	tbPassword.SetMaxLength(10)

	obj = GetObject("btnRegistration", builder)
	btnRegistration := obj.(*gtk.Button)


	btnRegistration.Connect("clicked", func() {
		login,_:= tbLogin.GetText()
		password, _:= tbPassword.GetText()
		err := clientHTTP.Registration(login,password)
		if err!=nil{
			ShowMessage("Ошибка создания пользователя!", win)
			return
		}
		ShowMessage("Вы успешно зарегистрировались!", win)
			win.Destroy()
		//
	})
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.ShowAll()
	gtk.Main()

}
