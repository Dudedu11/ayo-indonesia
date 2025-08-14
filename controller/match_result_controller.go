package controller

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchResultController struct {
	service service.MatchResultService
}

func NewMatchResultController(s service.MatchResultService) *MatchResultController {
	return &MatchResultController{service: s}
}

func (mrc *MatchResultController) Create(c *gin.Context) {
	var req dto.CreateMatchResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "data": nil})
		return
	}

	result, err := mrc.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "MatchResult created successfully", "data": result})
}

func (mrc *MatchResultController) GetAll(c *gin.Context) {
	results, err := mrc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": results})
}

func (mrc *MatchResultController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID", "data": nil})
		return
	}

	result, err := mrc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": result})
}

func (mrc *MatchResultController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid match result ID", "data": nil})
		return
	}

	var req dto.CreateMatchResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}

	result, err := mrc.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "MatchResult updated successfully", "data": result})
}

func (mrc *MatchResultController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid match result ID", "data": nil})
		return
	}

	if err := mrc.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "MatchResult deleted successfully", "data": []any{}})
}
