package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculler/techtuesdayapi/domain"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService domain.IUserService
}

func (h UserHandler) HandleUserGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdParam := ctx.Param("userId")

		userId, err := strconv.Atoi(userIdParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid userId in path")
			return
		}

		user, apiErr := h.UserService.GetById(userId)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (h UserHandler) HandlerUserGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, apiErr := h.UserService.GetAll()
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (h UserHandler) HandleUserCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *domain.User
		ctx.BindJSON(&user)

		resp, apiErr := h.UserService.Create(user)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusCreated, resp)
	}
}

func (h UserHandler) HandleUserUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *domain.User
		ctx.BindJSON(&user)

		_, apiErr := h.UserService.Update(user)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h UserHandler) HandleUserDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdParam := ctx.Param("userId")

		userId, err := strconv.Atoi(userIdParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid userId in path")
			return
		}

		_, apiErr := h.UserService.Delete(userId)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}