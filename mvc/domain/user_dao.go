package domain

import (
	"fmt"
	"net/http"

	"github.com/shraddhagajul/golang-microservices/mvc/utils"
)


var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Shraddha", LastName: "Gajul", Email: "shrads277@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	
	if user := users[userId] ;user != nil {
		return user,nil
		
	}
	return nil, &utils.ApplicationError{
		Message : fmt.Sprintf("user %v does not exists", userId),
		StatusCode : http.StatusNotFound,
		Code : "not_found",
	}
}