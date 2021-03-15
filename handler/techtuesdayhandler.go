package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sculler/techtuesdayapi/domain"
	"net/http"
	"strconv"
)

type TechTuesdayHandler struct {
	TechTuesdayService domain.ITechTuesdayService
}

const invalidUpdateParam = "invalid update path parameter"

func (h TechTuesdayHandler) HandleTechTuesdayGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		techTuesdayIdParam := ctx.Param("techTuesdayId")

		techTuesdayId, err := strconv.Atoi(techTuesdayIdParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid techTuesdayId in path")
			return
		}

		techTuesday, apiErr := h.TechTuesdayService.GetById(techTuesdayId)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusOK, techTuesday)
	}
}

func (h TechTuesdayHandler) HandleTechTuesdayGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		techTuesdays, apiErr := h.TechTuesdayService.GetAll()
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusOK, techTuesdays)
	}
}

func (h TechTuesdayHandler) HandleTechTuesdayCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var techTuesday *domain.TechTuesday
		ctx.BindJSON(&techTuesday)

		resp, apiErr := h.TechTuesdayService.Create(techTuesday)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusCreated, resp)
	}
}

func (h TechTuesdayHandler) HandleTechTuesdayUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var techTuesday *domain.TechTuesday
		ctx.BindJSON(&techTuesday)

		techTuesdayIdParam := ctx.Param("techTuesdayId")

		techTuesdayId, err := strconv.Atoi(techTuesdayIdParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid techTuesdayId in path")
			return
		}

		if techTuesday.ID != uint(techTuesdayId) {
			ctx.JSON(http.StatusBadRequest, invalidUpdateParam)
			return
		}

		_, apiErr := h.TechTuesdayService.Update(techTuesday)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h TechTuesdayHandler) HandleTechTuesdayDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		techTuesdayIdParam := ctx.Param("techTuesdayId")

		techTuesdayId, err := strconv.Atoi(techTuesdayIdParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid techTuesdayId in path")
			return
		}

		_, apiErr := h.TechTuesdayService.Delete(techTuesdayId)
		if apiErr != nil {
			ctx.JSON(apiErr.Code, apiErr.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}