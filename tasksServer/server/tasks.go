package server

import (
	"../storage"
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"time"

	"log"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	getTasks:= GetTasksRequest{}
	err = json.Unmarshal(body, &getTasks)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}

	user, err:= storage.GetUser(getTasks.Login)
	if err!=nil{
		http.Error(w, "Ошибка получения пользователя", 422)
		return
	}
	messagesSlice := StringToSlice(user.Tasks.String)

	tasksResp :=GetTasksResponse{}

	for _, value:= range messagesSlice{
		task, _:= storage.GetTask(value)


		if task.IdOrganisation.String == getTasks.IdOrganisation {

			tasksResp.Tasks = append(tasksResp.Tasks, task)
		}
	}
	jsData, err:= json.Marshal(tasksResp)
	if err!=nil{
		http.Error(w, "Ошибка получения задач", 422)
		return
	}

	w.Write(jsData)
}

func CreateTask(w http.ResponseWriter, r * http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	createTask:= CreateTasksRequest{}
	err = json.Unmarshal(body, &createTask)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}
	dialog:= storage.Dialog{}
	dialog.Id = createTask.Login + GetDateString()
	dialog.Name.String = "Задача:" +createTask.Name.String
	dialog.Name.Valid = true

	dialog.UsersId = createTask.Login + ";"
	dialog.IdCreator.String = createTask.Login
	dialog.IdCreator.Valid = true
	err = storage.CreateDialog(dialog)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка создания диалога", 500)
		return
	}
	task:= storage.Task{}
	task.Id = createTask.Login+GetDateString()
	task.IdOrganisation = sql.NullString{createTask.IdOrganisation, true}
	task.IdDialog = sql.NullString{dialog.Id, true}
	task.DateStart = mysql.NullTime{time.Now(), true}
	task.DateFinish = createTask.DateFinish
	task.Name = createTask.Name
	task.CreatorId = createTask.Login
	task.Status = 0
	task.Description = createTask.Description

	err = storage.AddTask(task)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка создания задачи", 500)
		return
	}
	err = storage.AddTaskToUser(task.Id,task.CreatorId)

	if err!=nil{
		storage.DeleteTask(task.Id)
		http.Error(w, "Добавления задачи к пользователю", 500)
		return
	}
	}

func GetTask(w http.ResponseWriter, r *http.Request){
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	getTasks:= GetTaskRequest{}
	err = json.Unmarshal(body, &getTasks)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}
	task, err:= storage.GetTask(getTasks.IdTask)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка загрузки задачи", 422)
		return
	}
	usersTask:= StringToSlice(task.UsersId)
	if !CheckUserInDialog(usersTask,getTasks.Login){
		log.Println(err)
		http.Error(w, "Ошибка доступа к задаче", 422)
		return
	}
	jsData, err:= json.Marshal(task)
	if err!=nil{
		http.Error(w, "Ошибка получения задач", 422)
		return
	}

	w.Write(jsData)
}

func DeleteUserFromTask(w http.ResponseWriter, r *http.Request)  {
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	editTask:= EditTaskRequest{}
	err = json.Unmarshal(body, &editTask)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}
	task, err:= storage.GetTask(editTask.TaskId)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка получения задачи", 422)
		return
	}
	if task.CreatorId!=editTask.Login{
		log.Println(err)
		http.Error(w, "Ошибка доступа", 403)
		return
	}
	if !CheckUserInDialog(StringToSlice(task.UsersId),editTask.UserLogin){
		http.Error(w, "Пользователя нет в задаче!", 422)
		return
	}
	err = storage.DeleteUserFromTask(editTask.UserLogin,task.Id)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка добавления", 403)
		return
	}
	err = storage.DeleteTaskFromUser(editTask.UserLogin,task.Id)
	if err!=nil{
		storage.AddUserToTask(editTask.UserLogin,editTask.TaskId)
		log.Println(err)
		http.Error(w, "Ошибка добавления", 403)
		return
	}
}

func AddUserToTask(w http.ResponseWriter, r *http.Request)  {
	body, err, code:= CheckGeneralRequest(http.MethodPost,w,r)
	if err!=nil{
		http.Error(w, err.Error(),code)
		return
	}
	editTask:= EditTaskRequest{}
	err = json.Unmarshal(body, &editTask)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка преобразования тела запроса", 422)
		return
	}
	task, err:= storage.GetTask(editTask.TaskId)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка получения задачи", 422)
		return
	}
	_, err = storage.GetUser(editTask.UserLogin)
	if err!=nil{
		if err!=nil{
			http.Error(w, "Данного пользователя нет в базе!", 422)
			return
		}
	}
	if task.CreatorId!=editTask.Login{
		log.Println(err)
		http.Error(w, "Ошибка доступа", 403)
		return
	}
	if CheckUserInDialog(StringToSlice(task.UsersId),editTask.UserLogin){
		http.Error(w, "Пользователь уже есть в задаче!", 403)
		return
	}
	err = storage.AddTaskToUser(task.Id,editTask.UserLogin)
	if err!=nil{
		log.Println(err)
		http.Error(w, "Ошибка добавления", 403)
		return
	}
	err = storage.AddUserToTask(editTask.UserLogin,task.Id)
	if err!=nil{
		storage.DeleteTaskFromUser(task.Id, editTask.UserLogin)
		log.Println(err)
		http.Error(w, "Ошибка добавления", 403)
		return
	}
}