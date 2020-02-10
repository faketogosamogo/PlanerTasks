package GUI

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func GetBuilderForFile(filename string)(*gtk.Builder){
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("GetBuilder: ", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("./gtkForms/"+filename)
	if err != nil {
		log.Fatal("GetBuilder:", err)
	}
	return b
}
func GetObject(nameobject string, builder *gtk.Builder)(glib.IObject){
	obj, err:= builder.GetObject(nameobject)

	if err!=nil{
		log.Fatal("Get object:", err)
	}
	return obj
}
func ShowMessage(message string, win *gtk.Window){
	mes := gtk.MessageDialogNew(win,gtk.DIALOG_MODAL,gtk.MESSAGE_INFO,gtk.BUTTONS_CLOSE,message)

	mes.Run()
	mes.Destroy()
}
func CreateColumn(title string, id int) *gtk.TreeViewColumn {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("Unable to create text cell renderer:", err)
	}

	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("Unable to create cell column:", err)
	}

	return column
}
