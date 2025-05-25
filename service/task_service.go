package service

import (
	"encoding/json"
	"errors"
	"strings"
	"todo-list-cli/dto"
	"todo-list-cli/model"
	"todo-list-cli/utils"
)

var filePath = "data/todo.json" // JSON file path

// function GetAllTasks to read all task from JSON file
func GetAllTasks() ([]model.Task, error) {
	data, err := utils.ReadFile(filePath) // read the file
	if err != nil {
		return nil, err
	}
	var tasks []model.Task
	json.Unmarshal(data, &tasks) // parsing JSON to Task slice
	return tasks, nil
}

// function SaveAllTasks to save the task into JSON FILE
func SaveAllTasks(tasks []model.Task) error {
	return utils.WriteFile(filePath, tasks)
}

// function AddTask to add new task with validation
func AddTask(req dto.CreateTaskRequest) error {
	if req.Title == "" {
		return errors.New("judul tugas tidak boleh kosong") // show error if input is empty
	}
	tasks, _ := GetAllTasks()
	for _, t := range tasks {
		if strings.EqualFold(t.Title, req.Title) { //check if task title is duplicate
			return errors.New("judul tugas sudah ada")
		}
	}
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1 // auto-increment ID based on the task ID
	}
	tasks = append(tasks, model.Task{ID: newID, Title: req.Title, Description: req.Description}) // add the task
	return SaveAllTasks(tasks)
}

// function MarkDone to mark the task is done based on task ID
func MarkDone(id int) error {
	tasks, _ := GetAllTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true // change the status to true
			return SaveAllTasks(tasks)
		}
	}
	return errors.New("tugas tidak ditemukan") // show error if Task ID is not found
}

// function DeleteTask to delete the task based on task ID
func DeleteTask(id int) error {
	tasks, _ := GetAllTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // delete the task slice element based on index
			return SaveAllTasks(tasks)
		}
	}
	return errors.New("tugas tidak ditemukan") // show error if Task ID is not found
}

// SearchTask to search the task based on title or description keyword
func SearchTask(req dto.SearchTaskRequest) []model.Task {
	var results []model.Task
	if strings.TrimSpace(req.Keyword) == "" {
		return results // if keyword is empty, return empty slice
	}
	tasks, _ := GetAllTasks()
	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), strings.ToLower(req.Keyword)) ||
			strings.Contains(strings.ToLower(t.Description), strings.ToLower(req.Keyword)) {
			results = append(results, t) // add the task which contain the keyword
		}
	}
	return results
}
