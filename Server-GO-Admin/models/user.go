package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

func (user *User) SetPassword(password string) {

	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		// Handle error, misalnya dengan logging atau propagasi error
		log.Printf("Error hashing password: %v", err)
		return
	}
	user.Password = newPassword

}

func (user *User) ComparePassword(password string) error {
	// Membandingkan password yang diberikan dengan hashed password yang tersimpan
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
