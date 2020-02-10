package GUI

import (
	"../clientHTTP"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"time"
)
var idMessage string
func SetTreeMessages(treeDialogs *gtk.TreeView){
	treeDialogs.AppendColumn(CreateColumn("Sender",0))
	treeDialogs.AppendColumn(CreateColumn("Text",1))
	treeDialogs.AppendColumn(CreateColumn("TimeSend",2))
}
func SetListMessages(listDialogs *gtk.ListStore){
	listDialogs.Clear()
	dialogs, err:= clientHTTP.GetMessages(Login,idDialog,Token)

	if err!=nil{
		log.Fatal(err)
	}
	for _, v:= range dialogs{
		iter:= listDialogs.Append()

		listDialogs.Set(iter,[]int{0,1,2},[]interface{}{v.LoginSender,v.Text,v.Time.Time.String()})

	}
}
func ShowDialogWindow(){
	gtk.Init(nil)
	builder:= GetBuilderForFile("MessagesWindow.glade")

	obj := GetObject("MessagesWindow",builder)
	win:= obj.(*gtk.Window)

	obj = GetObject("treeMessages", builder)
	treeMessages:= obj.(*gtk.TreeView)
	obj = GetObject("scrolledwindow1", builder)
	scrolledWindow:= obj.(*gtk.ScrolledWindow)

	obj = GetObject("btnSendFile", builder)
	btnSendFile:= obj.(*gtk.Button)
	obj = GetObject("filechooserbutton1", builder)
	fileChooser:= obj.(*gtk.FileChooserButton)

	obj = GetObject("btnSendMessage", builder)
	btnSendMessage:= obj.(*gtk.Button)
	obj = GetObject("entryMessage", builder)
	entryMessage:= obj.(*gtk.Entry)

	listMessages, _:= gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING,glib.TYPE_STRING)

	treeMessages.SetModel(listMessages)
	SetTreeMessages(treeMessages)


	go func() {
		for{
			adj:= scrolledWindow.GetVAdjustment()

			_, err:= glib.IdleAdd(SetListMessages,listMessages)
			if err!=nil{
				log.Fatal(err)
			}
			adj.SetValue(adj.GetUpper() - adj.GetPageSize())
			scrolledWindow.SetVAdjustment(adj)
			time.Sleep(time.Second * 5)

		}
	}()

	selMessages,_:= treeMessages.GetSelection()
	selMessages.SetMode(gtk.SELECTION_SINGLE)


	btnSendFile.Connect("clicked", func() {
		if fileChooser.GetFilename()!=""{
			fileName, err:= clientHTTP.UploadFile(Login,fileChooser.GetFilename(), Token)
			if err!=nil{
				log.Println(err)
				ShowMessage("Ошибка загрузки файла!", win)
				return
			}
			err = clientHTTP.SendMessage(Login,idDialog,Token,clientHTTP.FileGetURL+fileName,"file")
			if err!=nil{
				ShowMessage("Ошибка загрузки файла!", win)
				return
			}
		}
	})
	treeMessages.Connect("size-allocate", func(){

	})
	btnSendMessage.Connect("clicked", func() {
		text,_:= entryMessage.GetText()
		if len(text)==0{
			return
		}
		err := clientHTTP.SendMessage(Login,idDialog,Token,text, "text")
		if err!=nil{
			ShowMessage("Ошибка отправки сообщения", win)
			return
		}
		SetListMessages(listMessages)
	})

	win.Connect("destroy", func(){
		win.Destroy()
		gtk.MainQuit()
		ShowMainWindow()
	})
	win.ShowAll()
	adj:= scrolledWindow.GetVAdjustment()
	adj.SetValue(adj.GetLower())
	scrolledWindow.SetVAdjustment(adj)
	gtk.Main()
}
