package controllers

import (
	"sample_rest_api/parameters"
	"sample_rest_api/repositories"

	"github.com/kataras/iris/v12"
)

type userController struct {
	UserRepository    repositories.UserRepository
	SessionRepository repositories.SessionRepository
}

type UserController interface {
	CreateUser(ctx iris.Context)
	LoginUser(ctx iris.Context)
}

func NewUserController(ur repositories.UserRepository, sr repositories.SessionRepository) UserController {
	return &userController{UserRepository: ur, SessionRepository: sr}
}

func (c userController) CreateUser(ctx iris.Context) {
	var params parameters.NewUserParams
	err := ctx.ReadBody(&params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	user, err := c.UserRepository.CreateUser(params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, user)
}
func (c userController) LoginUser(ctx iris.Context) {
	var params parameters.LoginParams
	err := ctx.ReadBody(&params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	userID, err := c.UserRepository.Login(params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	session, err := c.SessionRepository.GenerateSession(parameters.NewSessionParams{UserID: userID})
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, session)

}
