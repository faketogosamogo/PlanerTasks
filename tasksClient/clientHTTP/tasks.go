package clientHTTP

import (
	"../storage"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
func GetTask(login, idTask, token string) (storage.Task, error){
	task:= storage.Task{}
	getTask:= GetTaskRequest{}
	getTask.Login = login
	getTask.IdTask = idTask

	js, err := json.Marshal(getTask)
	if err!=nil{
		return task, err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)


	req, err:= http.NewRequest("POST", GetTaskURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil{
		return task, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(body, &task)

	return task, err
}

func GetTasks(login,idOrganisation, token string) ([]storage.Task, error){
	getTasks:=GetTasksRequest{}
	getTasks.Login = login
	getTasks.IdOrganisation = idOrganisation


	js, err := json.Marshal(getTasks)
	if err!=nil{
		return nil, err
	}
	buf:= bytes.Buffer{}
	buf.Write(js)


	req, err:= http.NewRequest("POST", GetTasksURL,&buf)
	req.Header.Set("Access-Token", token)
	if err!=nil{
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, err
	}
	tasks:= GetTasksResponse{}
	err = json.Unmarshal(body, &tasks)
	return tasks.Tasks, err
}




