package entity

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "deepthinking.do/usersrolesapps/database"
)

type Role struct {
	Id       int    `json:"Id"`
	Rolename string `json:"RoleName"`
	Roledesc string `json:"RoleDesc"`
}

//Send All Roles
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()

	var roles []Role
	db.Find(&roles)

	fmt.Println("EndPoint Hit: All Roles Endpoint")
	json.NewEncoder(w).Encode(roles)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Role POST Endpoint Worked")
}
