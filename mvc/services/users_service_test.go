package services

import (
	"net/http"
	"testing"

	"github.com/shraddhagajul/golang-microservices/mvc/domain"
	"github.com/shraddhagajul/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var(
	userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func init(){
	 domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return getUserFunction(userId)
}
func TestGetUserNotFoundInDatabase(t *testing.T){
	getUserFunction = func(userId int64)(*domain.User, *utils.ApplicationError){
		return nil,&utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 does not exists",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t,user)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusNotFound,err.StatusCode)
	assert.EqualValues(t,"user 0 does not exists",err.Message)
}