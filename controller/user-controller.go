package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tai9/golang_jwt/dto"
	"github.com/tai9/golang_jwt/helper"
	"github.com/tai9/golang_jwt/service"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService,
		jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var userToUpdate dto.UserUpdateDTO
	err := ctx.ShouldBind(&userToUpdate)
	if err != nil {
		res := helper.BuildErrResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, errString := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if errString != nil {
		panic(err.Error())
	}
	userToUpdate.ID = id
	u := c.userService.Update(userToUpdate)
	res := helper.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}
func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	u := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}
