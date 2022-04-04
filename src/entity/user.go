package entity

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "deepthinking.do/usersrolesapps/database"
	handler "deepthinking.do/usersrolesapps/handler"
	"github.com/gorilla/mux"
	//"github.com/gin-gonic/gin"
)

type User struct {
	Id       uint64 `json:"Id"`
	Username string `json:"UserName"`
	Password string `json:"Password"`
}

//Login function
func Login(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()

	login := User{} //user in the database
	user := User{}  //user from json

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	//db.Find(&users).Where("username = ? and password = ?", user.Username, user.Password)
	if err := db.Raw("select username, password from users where username = ? and password = unhex(md5(?));", user.Username, user.Password).Scan(&login).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	} else if login.Username == "" {
		handler.RespondError(w, http.StatusUnauthorized, "Auth Failed!!!")
		return
	}

	handler.RespondMessageToJson(w, http.StatusOK, "Auth Successful!!!")
}

//Send All Users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// users := Users{
	// 	User{Id: 1, UserName: "Peter", Password: "Dan"},
	// 	User{Id: 2, UserName: "Di", Password: "Polo"},
	// 	User{Id: 3, UserName: "Vingo", Password: "Benga"},
	// 	User{Id: 4, UserName: "Teter", Password: "tenten"},
	// }

	// dsn := "dextro:1234@tcp(192.168.100.19:3306)/auth"
	// //dsn := "localhost:1234@tcp(192.168.100.19:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//Retrieve all records.
	//db.Find(&Users{})

	db := db.GetConnection()

	var users []User
	db.Find(&users)

	fmt.Println("EndPoint Hit: All Users Endpoint")

	json.NewEncoder(w).Encode(users)
	//json.NewEncoder(c.Writer).Encode(users)
}

//Send specific user.
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	dbuser := User{}
	vars := mux.Vars(r)

	v := vars["username"]

	db.Table("users").Where("username = ?", v).First(&dbuser)
	fmt.Println("EndPoint Hit: All Users Endpoint")

	json.NewEncoder(w).Encode(dbuser)
	//json.NewEncoder(c.Writer).Encode(users)
}

//Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User POST Endpoint Worked")

	db := db.GetConnection()
	user := User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Raw("insert into users(username, password) values(?,unhex(md5(?)));", user.Username, user.Password).Scan(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, user)
}

//Update current user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User PUT Endpoint Worked")

	db := db.GetConnection()
	user := User{}   //user in json
	dbuser := User{} //user in db

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	db.Table("users").Where("username = ?", user.Username).First(&dbuser)

	dbuser.Password = user.Password

	if err := db.Table("users").Where("username = ?", user.Username).Updates(&dbuser).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RespondJSON(w, http.StatusCreated, dbuser)
}

//Delete specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User DELETE Endpoint Worked")

	db := db.GetConnection()
	user := User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Where("username = ?", user.Username).Delete(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RespondJSON(w, http.StatusCreated, user)
}
