package storage

type File struct{
	Id string
	Type string
	Name string
	IdDIalog string
}

func AddFile(f File)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("insert into files values(?,?,?,?)",f.Id,f.Type, f.Name, f.IdDIalog)

	return err
}
func GetFile(id string)(File,error){
	db:=dbOpen()
	defer db.Close()
	file:= File{}
	err:= db.QueryRow("select * from files where id=?", id).Scan(&file.Id, &file.Type, &file.Name, &file.IdDIalog)
	return file, err
}
func DeleteFile(id string)error{
	db:=dbOpen()
	defer db.Close()

	_, err:= db.Exec("delete from files where id=?",id)

	return err
}