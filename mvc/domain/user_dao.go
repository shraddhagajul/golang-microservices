package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shraddhagajul/golang-microservices/mvc/utils"
)


var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Shraddha", LastName: "Gajul", Email: "shrads277@gmail.com"},
	}
	// UserDao userDao
	UserDao userDaoInterface
)

func init(){
	 UserDao = &userDao{}
	 	 UserDao = &userDao1{}

}
type userDaoInterface interface{
	GetUser(userId int64) (*User, *utils.ApplicationError)
}

type userDao struct{}
type userDao1 struct{}

func (u *userDao)GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing the database.")
	if user := users[userId] ;user != nil {
		return user,nil
		
	}
	
	return nil, &utils.ApplicationError{
		Message : fmt.Sprintf("user %v does not exists", userId),
		StatusCode : http.StatusNotFound,
		Code : "not_found",
	}
}

func (u *userDao1)GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing the database.")
	if user := users[userId] ;user != nil {
		return user,nil
		
	}
	return nil, &utils.ApplicationError{
		Message : fmt.Sprintf("user %v does not exists", userId),
		StatusCode : http.StatusNotFound,
		Code : "not_found",
	}
}