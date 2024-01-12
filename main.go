package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)

	// loop throughh courses and select the course and return that course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("There is no course with this id !")
	return
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course !")

	w.Header().Set("Content-Type", "application/json")

	// What if no data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter Some Data")
		return
	}
	// What if data like {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {

		json.NewEncoder(w).Encode("you send an empty json")
		return
	}

	// generate random id and add course to the courses

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return

}

// update data by id

func updateDataById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("updating data")

	w.Header().Set("Content-Type", "aplication/json")

	// retreive id from user

	// params := mux.Vars(r)

	//loop, id, remove, add value again in this fake DB

}

// Handlers - file

func main() {

	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/getall", GetAllData)
	http.HandleFunc("/getid", GetDatabyId)

	http.ListenAndServe(":8090", nil)

}
