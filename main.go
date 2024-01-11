package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for Course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// Model for Author - file
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Controller - file
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is an api</h1>"))
}

func GetAllData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is the data: ")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetDatabyId(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is the data by id: ")

	w.Header().Set("Content-Type", "application/json")

	// grab id from courses
	// params := mux.Vars(r)

	// loop throughh courses and select the course and return that course
	// for _, course := range courses {
	// 	if course.CourseId == params["id"] {

	// 	}
	// }

}

// Handlers - file

func main() {

	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/getall", GetAllData)
	http.HandleFunc("/getid", GetDatabyId)

	http.ListenAndServe(":8090", nil)

}
