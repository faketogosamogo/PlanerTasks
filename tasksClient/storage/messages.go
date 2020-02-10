package storage

import "github.com/go-sql-driver/mysql"

type Message struct{
	Id string
	LoginSender string
	Text string
	Time mysql.NullTime
	DialogId string
	Type string

}

