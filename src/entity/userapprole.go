package entity

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "deepthinking.do/usersrolesapps/database"
	"gorm.io/gorm/clause"
)

//Implementing Tabler interface
// type Tabler interface {
// 	TableName() string
// }

func (UserAppRole) TableName() string {
	return "usersappsroles"
}

type UserAppRole struct {
	Id        int     `json:"Id"`
	Idapprole int     `json:"IdApp"`
	Iduser    int     `json:"IdUser"`
	AppRole   AppRole `gorm:"foreignKey:Idapprole"`
	User      User    `gorm:"foreignKey:Iduser"`
}

func GetAllUsersAppsRoles(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()

	//var usersappsroles []map[string]interface{}
	var usersappsroles []UserAppRole
	db.
		Preload("AppRole.App").
		Preload("AppRole.Role").
		Preload(clause.Associations).
		Table("usersappsroles").
		Find(&usersappsroles)

	//Another kinds to implement the same is showed below
	//db.Preload("User").Preload("AppRole").Table("usersappsroles").Find(&usersappsroles)

	// var apps []App
	// db.Find(&apps)

	// var roles []Role
	// db.Find(&roles)

	fmt.Println("EndPoint Hit: All UsersAppsRoles Endpoint")
	json.NewEncoder(w).Encode(usersappsroles)
}
