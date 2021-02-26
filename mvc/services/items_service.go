package services

import (
	"github.com/shraddhagajul/golang-microservices/mvc/domain"
	"github.com/shraddhagajul/golang-microservices/mvc/utils"
)

type itemsService struct{}
var(
	ItemsService itemsService
)
func (i *itemsService)GetUser(userId int64) (*domain.User, *utils.ApplicationError) {

	return domain.UserDao.GetUser(userId)
}
