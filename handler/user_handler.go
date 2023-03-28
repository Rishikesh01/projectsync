package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"projectsync/dto"
	"projectsync/service"
)

type UserHandler struct {
	svc     service.UserService
	svcAuth service.AuthService
}

func NewUserHandler(svc service.UserService, svcAuth service.AuthService) *UserHandler {
	return &UserHandler{svc: svc, svcAuth: svcAuth}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var signup dto.Register
	if err := ctx.ShouldBindJSON(&signup); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if err := u.svc.Register(signup); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	var update dto.UpdateUser
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if err := u.svc.UpdateUser(update); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	uid, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if err := u.svc.DeleteUser(uid); err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var sign dto.SignIn

	if err := ctx.ShouldBindJSON(&sign); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	token, err := u.svcAuth.AuthenticateUser(sign)
	if err != nil {
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
