package task

import (
	"encoding/json"
	"mime"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type taskServer struct {
	storage *TaskStorage
}

func NewTaskServer() *taskServer {
	storage := NewStorage()
	return &taskServer{storage: storage}
}

func (ts *taskServer) CreateTaskHandler(w http.ResponseWriter, req *http.Request) {
	type RequestTask struct {
		Text string `json:"text"`
	}

	type ResponseId struct {
		Id int64 `json:"id"`
	}

	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediaType != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	var requestTask RequestTask
	err = decoder.Decode(&requestTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.storage.CreateTask(requestTask.Text, false)
	renderJson(w, ResponseId{Id: id})
}

func (ts *taskServer) GetAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	tasks := ts.storage.GetAllTasks()
	renderJson(w, tasks)
}

func (ts *taskServer) GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task, err := ts.storage.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderJson(w, task)
}

func (ts *taskServer) DeleteAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	ts.storage.DeleteAllTasks()
}

func (ts *taskServer) DeleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ts.storage.DeleteTask(id)
}

func renderJson(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
