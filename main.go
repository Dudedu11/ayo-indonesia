package main

import (
	"ayo-indonesia/config"
	"ayo-indonesia/controller"
	"ayo-indonesia/model"
	"ayo-indonesia/repository"
	"ayo-indonesia/routes"
	"ayo-indonesia/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&model.Team{}, &model.Player{}, &model.Match{}, &model.GoalDetail{}, &model.MatchResult{})

	r := gin.Default()

	matchResultRepo := repository.NewMatchResultRepository()
	teamRepo := repository.NewTeamRepository()
	playerRepo := repository.NewPlayerRepository()
	matchRepo := repository.NewMatchRepository()
	goalDetailRepo := repository.NewGoalDetailRepository()

	matchResultService := service.NewMatchResultService(matchResultRepo)
	teamService := service.NewTeamService(teamRepo, matchResultRepo)
	playerService := service.NewPlayerService(playerRepo, goalDetailRepo)
	matchService := service.NewMatchService(matchRepo, teamRepo, matchResultRepo, goalDetailRepo, playerRepo)
	goalDetailService := service.NewGoalDetailService(goalDetailRepo, matchRepo, playerRepo, teamRepo, matchResultRepo)
	
	matchResultController := controller.NewMatchResultController(matchResultService)
	teamController := controller.NewTeamController(teamService)
	playerController := controller.NewPlayerController(playerService)
	matchController := controller.NewMatchController(matchService)
	goalDetailController := controller.NewGoalDetailController(goalDetailService)



	routes.SetupRoutes(r, teamController, playerController, matchController, goalDetailController, matchResultController)

	r.Run(":8080")
}
