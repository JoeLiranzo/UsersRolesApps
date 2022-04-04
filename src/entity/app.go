package entity

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "deepthinking.do/usersrolesapps/database"
	handler "deepthinking.do/usersrolesapps/handler"
	"github.com/gorilla/mux"
)

type App struct {
	Id      uint64 `json:"Id"`
	Appname string `json:"AppName"`
	Appdesc string `json:"AppDesc"`
}

//Send all apps.
func GetAllApps(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()

	var apps []App
	db.Find(&apps)

	fmt.Println("EndPoint Hit: All Apps Endpoint")
	json.NewEncoder(w).Encode(apps)
}

func GetApp(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	dbapp := App{}
	vars := mux.Vars(r)

	v := vars["appname"]

	db.Table("users").Where("username = ?", v).First(&dbapp)
	fmt.Println("EndPoint Hit: All Users Endpoint")

	json.NewEncoder(w).Encode(dbapp)
	//json.NewEncoder(c.Writer).Encode(users)
}

func CreateApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App POST Endpoint Worked")

	db := db.GetConnection()
	app := App{}

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&app).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, app)
}

//Update current user
func UpdateApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User PUT Endpoint Worked")

	db := db.GetConnection()
	app := App{}   //app in json
	dbapp := App{} //app in db

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	db.Table("apps").Where("id = ?", app.Id).First(&app)

	dbapp = app

	if err := db.Table("apps").Where("id = ?", app.Id).Updates(&dbapp).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RespondJSON(w, http.StatusCreated, dbapp)
}

//Delete specific user
func DeleteApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User DELETE Endpoint Worked")

	db := db.GetConnection()
	app := App{}

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Where("id = ?", app.Id).Delete(&app).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RespondJSON(w, http.StatusCreated, app)
}
