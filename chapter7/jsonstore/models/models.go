package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Orders []Order
	Data   string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Order struct {
	gorm.Model
	User User
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func (User) TableName() string {
	return "users"
}

func (Order) TableName() string {
	return "orders"
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "postgres://postgres:root@localhost:5432/chapter7?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if !db.HasTable("users") {
		db.CreateTable(&User{})
	}
	if !db.HasTable("orders") {
		db.CreateTable(&Order{})
	}

	db.AutoMigrate(&User{}, &Order{})
	return db, nil
}
