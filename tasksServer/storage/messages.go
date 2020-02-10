package storage

import (
	"github.com/go-sql-driver/mysql"
)

type Message struct{
	Id string
	LoginSender string
	Text string
	Time mysql.NullTime
	DialogId string
	Type string

}

func AddMessage(m Message)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("insert into messages values(?,?,?,?,?,?)",m.Id,m.LoginSender, m.Text, m.Time, m.DialogId, m.Type)

	return err
}
func GetMessage(id string)(Message,error){
	db:=dbOpen()
	defer db.Close()
	mes:= Message{}
	err:= db.QueryRow("select * from messages where id=?", id).Scan(&mes.Id,&mes.LoginSender,&mes.Text,&mes.Time, &mes.DialogId, &mes.Type)

	return mes, err
}
func DeleteMessage(id string)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("delete from messages where id=?",id)

	return err
}