package models

import (
	"golang.org/x/crypto/bcrypt"
	"red-coins/config"
	"time"
)

type User struct {
	ID        int    `json:"id" gorm:"primary_key;auto_increment"`
	Name 	  string `gorm:"not null;" validate:"required" json:"name"`
	Email     string `gorm:"unique;not null;" validate:"required,email" json:"email"`
	Password  string `gorm:"not null;" validate:"required" json:"password"`
	Born 	  time.Time `gorm:"not null;" json:"born"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	
	return nil
}

func UserList() []User {
	var users []User
	res := config.DB.Find(&users)
	
	if res.Error == nil {
		return users
	}
	
	return nil
}

func UserStore(user *User) bool {
	res := config.DB.Debug().Create(&user)
	
	if res.Error == nil {
		return true
	}
	
	return false
}

func UserShow(id int) *User {
	user := new(User)
	res := config.DB.First(&user, id)
	
	if res.Error == nil {
		return user
	}
	
	return nil
}

func UserUpdate(user *User) bool {
	res := config.DB.Save(&user)
	
	if res.Error == nil {
		return true
	}

	return false
}

func UserDelete(id int) bool {
	user := new(User)
	res := config.DB.Where("id = ?", id).Delete(&user)
	
	if res.Error == nil {
		return true
	}
	
	return false
}
