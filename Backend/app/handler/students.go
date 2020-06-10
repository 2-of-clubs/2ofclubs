package handler

import (
	"../model"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateStudent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Check if content type is application/json?
	s := model.NewStudent()
	p := ExtractPersonInfo(r)
	pass, hashed := GenerateSaltedPass(p.Password)
	p.Password = pass
	s.Person = p
	found := RecordExists(db, model.ColumnUsername, s.Username, s)
	if !found && hashed {
		db.Create(&s)
		if token, err  := GetTokenPair(s.Username, 5, 60*24); err == nil {
			WriteData(ParseJSON(token), http.StatusOK, w)
		}

	}
}

func GetStudent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars[model.ColumnUsername]
	s := model.NewStudent()
	ss := model.New()
	// Defaults will be overridden when obtaining data and being inserted into struct except for null
	found := RecordExists(db, model.ColumnUsername, username, s)
	if !found {
		ss.Message = model.FAILURE
	} else {
		ss.Data = s
	}
	WriteData(ParseJSON(ss), http.StatusOK, w)
}

func UpdateStudent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Student")
}
