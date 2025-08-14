package controller

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	service service.TeamService
}

func NewTeamController(s service.TeamService) *TeamController {
	return &TeamController{service: s}
}

func (tc *TeamController) Create(c *gin.Context) {
	var req dto.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "data": nil})
		return
	}
	team, err := tc.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Team created successfully", "data": team})
}

func (tc *TeamController) GetAll(c *gin.Context) {
	teams, err := tc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": teams})
}

func (tc *TeamController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID", "data": nil})
		return
	}
	team, err := tc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": team})
}

func (tc *TeamController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid team ID", "data": nil})
		return
	}
	var req dto.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	team, err := tc.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team updated successfully", "data": team})
}

func (tc *TeamController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid team ID", "data": nil})
		return
	}
	if err := tc.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted successfully", "data": []any{}})
}