package controller

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoalDetailController struct {
	service service.GoalDetailService
}

func NewGoalDetailController(s service.GoalDetailService) *GoalDetailController {
	return &GoalDetailController{service: s}
}

func (gc *GoalDetailController) Create(c *gin.Context) {
	var req dto.CreateGoalDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "data": nil})
		return
	}
	goal, err := gc.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Goal detail created successfully", "data": goal})
}

func (gc *GoalDetailController) GetAll(c *gin.Context) {
	goals, err := gc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": goals})
}

func (gc *GoalDetailController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID", "data": nil})
		return
	}
	goal, err := gc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": goal})
}

func (gc *GoalDetailController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid goal ID", "data": nil})
		return
	}
	var req dto.CreateGoalDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	goal, err := gc.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Goal detail updated successfully", "data": goal})
}

func (gc *GoalDetailController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid goal ID", "data": nil})
		return
	}
	if err := gc.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Goal detail deleted successfully", "data": []any{}})
}
