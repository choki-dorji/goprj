package routes

import (
	"log"
	contolller "myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	var port = 8080
	router := mux.NewRouter()

	router.HandleFunc("/student", contolller.AddStudent).Methods("POST")
	router.HandleFunc("/student", contolller.GetAllStuds).Methods("GET")
	router.HandleFunc("/student/{sid}", contolller.GetStdu).Methods("GET")
	router.HandleFunc("/student/{sid}", contolller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", contolller.DeleteUser).Methods("DELETE")

	router.HandleFunc("/enroll", contolller.Enroll).Methods("POST")
	router.HandleFunc("/enrolls", contolller.GetAllEnrolls).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", contolller.GetEnroll).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", contolller.DeleteEnroll).Methods("DELETE")
	router.HandleFunc("/enroll/{sid}", contolller.UpdateEnroll).Methods("PUT")

	//course
	router.HandleFunc("/course", contolller.AddCourse).Methods("POST")
	router.HandleFunc("/courses", contolller.GetCourses).Methods("GET")
	router.HandleFunc("/course/{cid}", contolller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", contolller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", contolller.DeleteCourse).Methods("DELETE")

	//admin
	router.HandleFunc("/signup", contolller.Signup).Methods("POST")
	router.HandleFunc("/login", contolller.Login).Methods("POST")
	router.HandleFunc("/logout", contolller.Logout)

	//serving static files
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("application runninng on the port", port)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
