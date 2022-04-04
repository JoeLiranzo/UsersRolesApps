package main

import (
	"fmt"
	"log"
	"net/http"

	db "deepthinking.do/usersrolesapps/database"
	d "deepthinking.do/usersrolesapps/deepthinking"
	e "deepthinking.do/usersrolesapps/entity"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage EndPoint Hit")
}

func handleRequests() {

	// router := gin.Default()

	// router.GET("/users", e.GetAllUsers)
	// router.Run()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/login", e.Login).Methods("GET")

	myRouter.HandleFunc("/users/{username}", e.GetUser).Methods("GET")
	myRouter.HandleFunc("/users/", e.GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users", e.CreateUser).Methods("POST")
	myRouter.HandleFunc("/users", e.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/users", e.DeleteUser).Methods("DELETE")

	myRouter.HandleFunc("/roles", e.GetAllRoles).Methods("GET")

	myRouter.HandleFunc("/apps/{appname}", e.GetApp).Methods("GET")
	myRouter.HandleFunc("/apps", e.GetAllApps).Methods("GET")
	myRouter.HandleFunc("/apps", e.CreateApp).Methods("POST")

	myRouter.HandleFunc("/appsroles", e.GetAllAppsRoles).Methods("GET")

	myRouter.HandleFunc("/usersappsroles", e.GetAllUsersAppsRoles).Methods("GET")

	myRouter.HandleFunc("/CPUFull", d.GetCPUFULL).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	db.Init()
	handleRequests()
}
