package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`

	Habits []Habit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName sets the table name for the User model.
func (User) TableName() string {
	return "users"
}

// SetPassword sets the password as a hash for the User model using bcrypt default behavior.
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the password and the hash for the User model using bcrypt default behavior.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
