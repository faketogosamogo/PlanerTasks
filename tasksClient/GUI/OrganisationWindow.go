package GUI

import (
	"../clientHTTP"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"time"
)

var idOrganisation string

func SetTreeOrganisations(treeOrganisations *gtk.TreeView){
	treeOrganisations.AppendColumn(CreateColumn("Id",0))
	treeOrganisations.AppendColumn(CreateColumn("Name",1))

}
func SetListOrganisations(listOrganisations *gtk.ListStore){
	listOrganisations.Clear()
	organisations, err:= clientHTTP.GetOrganisations(Login, Token)
	if err!=nil{
		log.Fatal(err)
	}

	for _, v:= range organisations{
		iter:= listOrganisations.Append()
		listOrganisations.Set(iter,[]int{0,1},[]interface{}{v.Id,v.Name})
	}
}
func ShowOrganisationWindow(){
	gtk.Init(nil)
	builder:= GetBuilderForFile("OrganisationWindow1.glade")

	obj:= GetObject("OrganisationWindow",builder)
	win := obj.(*gtk.Window)

	obj = GetObject("btnChoose", builder)
	btnChoose := obj.(*gtk.Button)

	obj = GetObject("treeOrganisations", builder)
	treeOrganisations:= obj.(*gtk.TreeView)

	listOrganisations, err:= gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err!=nil{
		log.Fatal(err)
	}
	SetTreeOrganisations(treeOrganisations)
	treeOrganisations.SetModel(listOrganisations)
	go func() {
		for{
			_, err:= glib.IdleAdd(SetListOrganisations,listOrganisations)
			if err!=nil{
				log.Fatal(err)
			}
			time.Sleep(time.Second * 5)
			}
	}()


	selOrganisations, _:= treeOrganisations.GetSelection()
	selOrganisations.SetMode(gtk.SELECTION_SINGLE)

	selOrganisations.Connect("changed", func() {

		rows := selOrganisations.GetSelectedRows(listOrganisations)

		if rows.Length()==0{
			return
		}
		items := make([]string, 0, rows.Length())

		for l := rows; l != nil; l = l.Next() {
			path := l.Data().(*gtk.TreePath)
			iter, _ := listOrganisations.GetIter(path)
			value, _ := listOrganisations.GetValue(iter, 0)
			str, _ := value.GetString()
			items = append(items, str)
		}

		idOrganisation = items[0]
	})
	btnChoose.Connect("clicked", func() {
		if idOrganisation !=""{
			win.SetVisible(false)
			win.Destroy()
			ShowMainWindow()
		}
	})
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})


	win.ShowAll()
	gtk.Main()
}
