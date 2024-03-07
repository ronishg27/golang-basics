package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for courses -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware or helper function - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func checkNilError(err error) {
	// if err != nil {
	// 	panic(err)
	// }

}
func main() {
	fmt.Println("Welcome to my API")

	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Course{"1", "ReactJS", 299, &Author{"Ronish", "github.com/ronishg27"}})
	courses = append(courses, Course{"2", "MERN Stack", 199, &Author{"Ronish", "github.com/ronishg27/MERN"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listening to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

// first response and then request
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API's homepage</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// loop through courses
	// find matching id
	// return response

	for _, course := range courses { //looping through courses
		if course.CourseId == params["id"] { //finding matching id
			json.NewEncoder(w).Encode(course) //returning response
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	checkNilError(err)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data sent")
		return
	}

	// check if courseid and coursename matches to existing course
	//  - loop, title & id, -> json

	// generating unique id
	// converting to string
	// append course into courses

	course.CourseId = strconv.Itoa(rand.Intn(10000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	// grabbing id from req
	params := mux.Vars(r)

	if r.Body == nil {
		json.NewEncoder(w).Encode("No data sent")
	}

	var recievedCourse Course
	err := json.NewDecoder(r.Body).Decode(&recievedCourse)
	checkNilError(err)

	// looping through values, id -> remove -> add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// var course Course
			// err := json.NewDecoder(r.Body).Decode(&course)
			// checkNilError(err)
			// course.CourseId = params["id"]
			// courses = append(courses, course)
			// json.NewEncoder(w).Encode(course)

			//
			recievedCourse.CourseId = params["id"]
			courses = append(courses, recievedCourse)
			json.NewEncoder(w).Encode(recievedCourse)
			//
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id , remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}
