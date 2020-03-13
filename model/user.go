package model

import (
	"errors"
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
)

var mySecret = []byte("My Secret")

type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserWrap struct {
	User
	Token string `json:"token"`
}

func (user *User) New() error {
	if !DB.Where("username = ?", user.Username).Find(&User{}).RecordNotFound() {
		return errors.New("user already exists")
	}
	return DB.Create(&user).Error
}

func (user *User) Login(userWrap *UserWrap) error {
	if DB.Where("username = ?", user.Username).First(&userWrap.User).RecordNotFound() {
		return errors.New("username does not exist")
	}
	fmt.Println("user", user, userWrap)
	if user.Password != userWrap.Password {
		return errors.New("incorrect password")
	}
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		return err
	}
	userWrap.Token = tokenString
	return nil
}

func GetAllUser(list *[]User) error {
	return DB.Order("id asc").Find(list).Error
}
