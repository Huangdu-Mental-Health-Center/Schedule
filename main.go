package main

import (
	"Huangdu_HMC_Schedule/src/handler"
	"Huangdu_HMC_Schedule/src/logger"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupLogs() {

	gin.DisableConsoleColor()
	f, err := os.Create("gin.log")
	if err != nil {
		logger.Error.Printf("create gin.log failed: %v;\n", err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
}

func main() {
	logger.Init()
	logger.Info.Println("Huangdu HMC Schedule module.")
	setupLogs()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	h := handler.Handler{}
	router.Use(cors.Default())
	router.GET("/api/schedule", h.GetDoctor)
	_ = router.Run(":5700")
}
