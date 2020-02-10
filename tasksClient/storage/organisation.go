package storage

import (

	"strings"
)

type Organisation struct{
	Id string
	IdUsers string
	IdDialogs string
	IdAdmins string
	Name string
	IdTasks string
}
func StringToSlice(str string)[]string{
	slice:= strings.Split(str,";")
	return slice
}
func SliceToString(slice []string)string{
	str:=""
	for _, value:= range slice{
		if value!="" {
			str += value + ";"
		}
	}
	return str
}