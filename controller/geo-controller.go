package controller

import (
	"fmt"
	"keruen-geo/helper"
	"keruen-geo/models"
	"keruen-geo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	GeoService service.GeoService
}

func NewController(service service.GeoService) controller {
	controller := controller{
		GeoService: service,
	}

	return controller
}

func (c *controller) Create(ctx *gin.Context) {
	var geo models.Location

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := ctx.ShouldBindJSON(&geo); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}
	// here service
	if err := c.GeoService.Create(id, geo); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}
	ctx.JSON(201, "ok")
}

func (c *controller) GetAllDrivers(ctx *gin.Context) {
	results, err := c.GeoService.GetAll()
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
	}

	ctx.JSON(200, results)
}

func (c *controller) GetDrivers(ctx *gin.Context) {
	result, err := c.GeoService.Get()

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	ctx.JSON(200, result)
}

func (c *controller) GetNearby(ctx *gin.Context) {

	var geo helper.Location

	if err := ctx.ShouldBindJSON(&geo); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	result, err := c.GeoService.GetNearby(geo)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	ctx.JSON(200, result)
}
