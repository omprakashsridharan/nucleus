package models

import (
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Model is sample of common table structure
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func init() {
	fmt.Println("Database Initialisation")
	username := "root"
	password := "12qwmn09"
	db := "nucleus"
	host := "127.0.0.1"
	port := "3306"

	msql := mysql.Config{}
	log.Println(msql)

	_, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+db+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")
	if err != nil {
		fmt.Print(err)
	}
	
}

//CreateSeed Test Data
func CreateSeed() {
	db.Create(&User{Name: "omprakash"})
}

//GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
