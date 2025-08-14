package routes

import (
	"ayo-indonesia/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	tc *controller.TeamController,
	pc *controller.PlayerController,
	mc *controller.MatchController,
	gdc *controller.GoalDetailController,
	mrc *controller.MatchResultController) {
	api := r.Group("/api")
	{
		api.POST("/teams", tc.Create)
		api.GET("/teams", tc.GetAll)
		api.GET("/teams/:id", tc.GetByID)
		api.PUT("/teams/:id", tc.Update)
		api.DELETE("/teams/:id", tc.Delete)

		api.POST("/players", pc.Create)
		api.GET("/players", pc.GetAll)
		api.GET("/players/:id", pc.GetByID)
		api.GET("/players/top-score", pc.GetTopScore)
		api.PUT("/players/:id", pc.Update)
		api.DELETE("/players/:id", pc.Delete)

		api.POST("/matchs", mc.Create)
		api.GET("/matchs", mc.GetAll)
		api.GET("/matchs/:id", mc.GetByID)
		api.PUT("/matchs/:id", mc.Update)
		api.DELETE("/matchs/:id", mc.Delete)

		api.POST("/goal-details", gdc.Create)
		api.GET("/goal-details", gdc.GetAll)
		api.GET("/goal-details/:id", gdc.GetByID)
		api.PUT("/goal-details/:id", gdc.Update)
		api.DELETE("/goal-details/:id", gdc.Delete)

		api.POST("/match-results", mrc.Create)
		api.GET("/matchs-results", mrc.GetAll)
		api.GET("/matchs-results/:id", mrc.GetByID)
		api.PUT("/matchs-results/:id", mrc.Update)
		api.DELETE("/matchs-results/:id", mrc.Delete)
	}
}
