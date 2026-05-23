package routes

import (
	"fmt"
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	// creating a new router
	router := mux.NewRouter()
	//register ROUTES ON mux router
	//student route
	router.HandleFunc("/student", controller.GetAllStuds).Methods("GET")
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")

	//Course Routes
	router.HandleFunc("/course", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/course", controller.GetAllCourses).Methods("GET")

	// signup and login
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/loginout", controller.Logout)

	// enroll APIs
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")
	//Load static files
	fhandler := http.FileServer(http.Dir("./view"))
	//server static files as a route by the registering all static files on the mux router
	router.PathPrefix("/").Handler(fhandler)

	fmt.Println("Server started successfully...")
	//start the http server
	log.Fatal(http.ListenAndServe(":8080", router))

}
