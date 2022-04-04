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

func (AppRole) TableName() string {
	return "appsroles"
}

type AppRole struct {
	Id     int  `json:"Id"`
	Idapp  int  `json:"IdApp"`
	Idrole int  `json:"IdRole"`
	App    App  `gorm:"foreignKey:Idapp"`
	Role   Role `gorm:"foreignKey:Idrole"`
}

func GetAllAppsRoles(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()

	var appsroles []AppRole

	db.
		Preload(clause.Associations).
		Table("appsroles").
		Find(&appsroles)

	//Another kinds to implement the same is showed below
	//db.Preload("App").Preload("Role").Table("appsroles").Find(&appsroles)

	// var app App
	// var role Role
	// //Adding foreign key content.
	// for idx, element := range appsroles {
	// 	db.Table("apps").Where("id = ?", element.Idapp).First(&app)
	// 	db.Table("roles").Where("id = ?", element.Idrole).First(&role)

	// 	element.App = app
	// 	element.Role = role
	// 	appsroles[idx] = element
	// }

	fmt.Println("EndPoint Hit: All AppsRoles Endpoint")
	json.NewEncoder(w).Encode(appsroles)
}
