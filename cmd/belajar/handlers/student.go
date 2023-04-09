package handlers

import (
	"belajar-go-orm/cmd/belajar/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func OutputJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return
}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

var students []*Student

func init() {
	students = []*Student{
		&Student{ID: 1, Name: "John Doe", Grade: 2},
		&Student{ID: 2, Name: "Jane Doe", Grade: 3},
	}

	students = append(students, &Student{ID: 3, Name: "John Doe", Grade: 2})
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id int) *Student {
	for _, student := range students {
		if student.ID == id {
			return student
		}
	}
	return nil
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(w, r) {
		return
	}
	if !middleware.OnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		fmt.Println("ID:", id)
		parseID, err := strconv.Atoi(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		OutputJSON(w, SelectStudent(parseID))
		return
	}

	OutputJSON(w, GetStudents())
}
