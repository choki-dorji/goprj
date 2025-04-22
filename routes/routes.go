package routes

import (
	"log"
	contolller "myapp/myapp/controller"
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

	//admin
	router.HandleFunc("/signup", contolller.Signup).Methods("POST")
	router.HandleFunc("/login", contolller.Login).Methods("POST")

	//serving static files
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("application runninng on the port", port)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
