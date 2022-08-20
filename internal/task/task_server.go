package task

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

type taskServer struct {
	storage *TaskStorage
}

func NewTaskServer() *taskServer {
	storage := NewStorage()
	return &taskServer{storage: storage}
}

func (ts *taskServer) Handler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/task/" {
		if req.Method == http.MethodPost {
			ts.createTaskHandler(w, req)
		} else if req.Method == http.MethodGet {
			ts.getAllTasksHandler(w, req)
		} else if req.Method == http.MethodDelete {
			ts.deleteAllTasksHandler(w, req)
		} else {
			http.Error(
				w,
				fmt.Sprintf("expect method GET, DELETE or POST at /task/, got %v", req.Method),
				http.StatusMethodNotAllowed,
			)
		}
	} else {
		trimPath := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(trimPath, "/")

		if len(pathParts) != 2 {
			http.Error(w, fmt.Sprintf("expect /task/{id}, got %v", req.URL.Path), http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(pathParts[1], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Method == http.MethodDelete {
			ts.deleteTaskHandler(w, req, id)
		} else if req.Method == http.MethodGet {
			ts.getTaskHandler(w, req, id)
		} else {
			http.Error(
				w,
				fmt.Sprintf("expect method GET or DELETE at /task/{id}, got %v", req.Method),
				http.StatusMethodNotAllowed,
			)
		}
	}
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Command: create_task. Path: %s\n", req.URL.Path)

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

	id := ts.storage.CreateTask(requestTask.Text)
	renderJson(w, ResponseId{Id: id})
}

func (ts *taskServer) getAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Command: get_all_task. Path: %s\n", req.URL.Path)

	tasks := ts.storage.GetAllTasks()
	renderJson(w, tasks)
}

func (ts *taskServer) getTaskHandler(w http.ResponseWriter, req *http.Request, id int64) {
	log.Printf("Command: get_task. Path: %s\n", req.URL.Path)

	task, err := ts.storage.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderJson(w, task)
}

func (ts *taskServer) deleteAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Command: delete_all_task. Path: %s\n", req.URL.Path)
	ts.storage.DeleteAllTasks()
}

func (ts *taskServer) deleteTaskHandler(w http.ResponseWriter, req *http.Request, id int64) {
	log.Printf("Command: delete_task. Path: %s\n", req.URL.Path)
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
