package controller

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	service service.PlayerService
}

func NewPlayerController(s service.PlayerService) *PlayerController {
	return &PlayerController{service: s}
}

func (pc *PlayerController) Create(c *gin.Context) {
	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	player, err := pc.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Player created successfully", "data": player})
}

func (pc *PlayerController) GetAll(c *gin.Context) {
	players, err := pc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": players})
}

func (pc *PlayerController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	player, err := pc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": player})
}

func (pc *PlayerController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	player, err := pc.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player updated successfully", "data": player})
}

func (pc *PlayerController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := pc.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully", "data": []any{}})
}

func (pc *PlayerController) GetTopScore(c *gin.Context) {
	players, err := pc.service.GetTopScore()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": players})
}