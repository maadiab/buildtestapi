package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func main() {

	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "go development",
		CoursePrice: 300, Author: &Author{Fullname: "mohanad diab", Website: "maadiab.io"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "reactjs",
		CoursePrice: 250, Author: &Author{Fullname: "mohanad ahmed", Website: "mohanad-diab.com"}})

	// Routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllData).Methods("GET")
	r.HandleFunc("/course/{id}", GetDatabyId).Methods("GET")
	r.HandleFunc("/course", CreateData).Methods("POST")
	r.HandleFunc("/course/{id}", updateDataById).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteSomeData).Methods("DELETE")

	// Listen to the port
	log.Fatal(http.ListenAndServe(":8090", r))
}

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

	params := mux.Vars(r)

	//loop, id, remove, add value again in this fake DB

	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteSomeData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete some data")

	w.Header().Set("Content-Type", "aplication/json")

	params := mux.Vars(r)

	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
