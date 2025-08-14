package controller

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchController struct {
	service service.MatchService
}

func NewMatchController(s service.MatchService) *MatchController {
	return &MatchController{service: s}
}

func (mc *MatchController) Create(c *gin.Context) {
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "data": nil})
		return
	}
	match, err := mc.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Match created successfully", "data": match})
}

func (mc *MatchController) GetAll(c *gin.Context) {
	matches, err := mc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": matches})
}

func (mc *MatchController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID", "data": nil})
		return
	}
	match, err := mc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": match})
}

func (mc *MatchController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid match ID", "data": nil})
		return
	}
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	match, err := mc.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Match updated successfully", "data": match})
}

func (mc *MatchController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid match ID", "data": nil})
		return
	}
	if err := mc.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Match deleted successfully", "data": []any{}})
}
