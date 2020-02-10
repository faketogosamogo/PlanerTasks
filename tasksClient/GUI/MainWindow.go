package GUI

import (
	"../clientHTTP"
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"time"

)
var idDialog string


var idTask string



func SetTreeDialogs(treeDialogs *gtk.TreeView){
	treeDialogs.AppendColumn(CreateColumn("Id",0))
	treeDialogs.AppendColumn(CreateColumn("Name",1))
}
func SetTreeTasks(treeTasks *gtk.TreeView){
	treeTasks.AppendColumn(CreateColumn("Id",0))
	treeTasks.AppendColumn(CreateColumn("Description",1))
}
func SetListDialogs(listDialogs *gtk.ListStore){
	listDialogs.Clear()
	dialogs, err := clientHTTP.GetDialogs(Login,idOrganisation,Token)

	if err!=nil{
		log.Fatal(err)
	}

	for _, v:= range dialogs{
		iter:= listDialogs.Append()
		listDialogs.Set(iter,[]int{0,1},[]interface{}{v.Id,v.Name.String})
	}
}
func SetListTasks(listTasks *gtk.ListStore){
	listTasks.Clear()
	err:= errors.New("")
	tasks, err:= clientHTTP.GetTasks(Login,idOrganisation,Token)

	if err!=nil{
		log.Fatal(err)
	}
	for _, v:= range tasks{
		iter:= listTasks.Append()
		listTasks.Set(iter,[]int{0,1},[]interface{}{v.Id,v.Description})
	}
}

func ShowMainWindow(){

	gtk.Init(nil)
	builder:= GetBuilderForFile("DialogsWindow.glade")

	obj := GetObject("DialogsWindow",builder)
	win:= obj.(*gtk.Window)

	obj = GetObject("btnCreateDialog", builder)
	btnCreateDialog:= obj.(*gtk.Button)

	obj = GetObject("entryNameDialog", builder)
	entryNameDialog:= obj.(*gtk.Entry)

	obj = GetObject("btnChooseDialog", builder)
	btnChooseDialog:= obj.(*gtk.Button)

	obj = GetObject("btnChooseTask", builder)
	btnChooseTask:= obj.(*gtk.Button)

	obj = GetObject("treeDialogs", builder)
	treeDialogs:= obj.(*gtk.TreeView)

	obj = GetObject("treeTasks", builder)
	treeTasks:= obj.(*gtk.TreeView)

	listDialogs, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	listTasks, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	treeDialogs.SetModel(listDialogs)
	SetTreeDialogs(treeDialogs)

	treeTasks.SetModel(listTasks)
	SetTreeTasks(treeTasks)



	SetListDialogs(listDialogs)
	SetListTasks(listTasks)
	go func() {
		for{
			_, err= glib.IdleAdd(SetListDialogs,listDialogs)

			if err!=nil{
				log.Fatal(err)
			}

			time.Sleep(time.Second * 5)

		}
	}()
	go func() {
		for{

			_, err:= glib.IdleAdd(SetListTasks,listTasks)
			if err!=nil{
				log.Fatal(err)
			}
			time.Sleep(time.Second * 5)

		}
	}()


	selDialogs, _:= treeDialogs.GetSelection()
	selDialogs.SetMode(gtk.SELECTION_SINGLE)

	selTasks,_:= treeTasks.GetSelection()
	selTasks.SetMode(gtk.SELECTION_SINGLE)
	btnChooseDialog.Connect("clicked", func() {
		win.Destroy()
		gtk.MainQuit()
		ShowDialogWindow()
	})

	btnChooseTask.Connect("clicked", func() {
		if idTask==""{
			return
		}
		win.Destroy()
		gtk.MainQuit()
		ShowTaskWindow()

	})
	win.Connect("destroy", func(){
		gtk.MainQuit()
	})

	btnCreateDialog.Connect("clicked", func() {
		nameDialog,_:= entryNameDialog.GetText()
		if len(nameDialog)<1{
			ShowMessage("Ошибка имени диалога!" ,win)
			return
		}
		err := clientHTTP.CreateDialog(Login,idOrganisation,nameDialog,Token)
		if err!=nil{
			ShowMessage("Ошибка создания диалога", win)
			return
		}
		ShowMessage("Диалог успешно создан!", win)
		SetListDialogs(listDialogs)
	})

	selTasks.Connect("changed", func() {
		rows := selTasks.GetSelectedRows(listTasks)

		if rows.Length()==0{
			return
		}
		items := make([]string, 0, rows.Length())

		for l := rows; l != nil; l = l.Next() {
			path := l.Data().(*gtk.TreePath)
			iter, _ := listTasks.GetIter(path)
			value, _ := listTasks.GetValue(iter, 0)
			str, _ := value.GetString()
			items = append(items, str)
		}

		idTask = items[0]

	})

	selDialogs.Connect("changed", func() {
		rows := selDialogs.GetSelectedRows(listDialogs)

		if rows.Length()==0{
			return
		}
		items := make([]string, 0, rows.Length())

		for l := rows; l != nil; l = l.Next() {
			path := l.Data().(*gtk.TreePath)
			iter, _ := listDialogs.GetIter(path)
			value, _ := listDialogs.GetValue(iter, 0)
			str, _ := value.GetString()
			items = append(items, str)
		}

		idDialog = items[0]
	})
	win.ShowAll()
	gtk.Main()
	}
