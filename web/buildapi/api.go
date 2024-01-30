package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// model for courses
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

func main() {
	fmt.Println("API DEMO IN GO LANG")

	//adding demo courses
	courses = append(courses, Course{
		CourseId:    "23542354235",
		CourseName:  "Sveltekit",
		CoursePrice: 50230,
		Author: &Author{
			Fullname: "sahil",
			Website:  "sahil.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "67565346",
		CourseName:  "react",
		CoursePrice: 503320,
		Author: &Author{
			Fullname: "dennis ritche",
			Website:  "dr.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "245635",
		CourseName:  "ML",
		CoursePrice: 5890,
		Author: &Author{
			Fullname: "samaltman",
			Website:  "samaltman.com",
		},
	})

	// initiating mux router
	router := mux.NewRouter()

	//routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createonecourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateonecourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteoneocurse).Methods("DELETE")

	//start the server
	log.Fatal(http.ListenAndServe(":4000", router))
}

// Isempty function
func (c *Course) Isempty() bool {
	return c.CourseName == ""
}

func serveHome(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("api home page"))
}

func getAllCourses(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("Get all courses")
	writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(writer).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")
	// check input data
	if r.Body == nil {
		json.NewEncoder(w).Encode("PLEASE GIVE SOME DATA")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty() {
		json.NewEncoder(w).Encode("course is empty")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateonecourse(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	writer.Header().Set("Content-Type", "application/json")
	//GRAB THAT COURSE USING ID
	params := mux.Vars(r)

	//loop over the course using id,remove,add updated with same id
	for index, course := range courses {

		if course.CourseId == params["id"] {
			//remove a value from slice using append
			courses = append(courses[:index], courses[index+1:]...)
			var coursetoadd Course
			_ = json.NewDecoder(r.Body).Decode(&coursetoadd)
			coursetoadd.CourseId = params["id"]
			courses = append(courses, coursetoadd)
			json.NewEncoder(writer).Encode(coursetoadd)
			return
		}

	}

	//when id is not found
	json.NewEncoder(writer).Encode("course with same i d not found")

}

func deleteoneocurse(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(writer).Encode("course deleted")
			return
		}
	}
	//if id not found
	json.NewEncoder(writer).Encode("no course found with such id")

}
