package main

import (
	"aditydcp/wfgo-web-service/controllers"
	config "aditydcp/wfgo-web-service/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/farm/:id", inDB.GetFarmById)
	router.GET("/farms", inDB.GetAllFarms)
	router.POST("/farm", inDB.CreateFarm)
	router.PUT("/farm", inDB.UpdateFarm)
	router.DELETE("/farm/:id", inDB.DeleteFarm)

	router.GET("/pond/:id", inDB.GetPondById)
	router.GET("/ponds", inDB.GetAllPonds)
	router.POST("/pond", inDB.CreatePond)
	router.PUT("/pond", inDB.UpdatePond)
	router.DELETE("/pond/:id", inDB.DeletePond)

	router.Run(":3000")
}
