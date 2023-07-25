package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/composer"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/postgre"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/setting"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.RouterGroup) {
	taskAPIService := composer.ComposerTaskAPIService(db)

	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskAPIService.GetTasks())
		tasks.GET("/:id", taskAPIService.GetTask())
		tasks.GET("", taskAPIService.CreateTask())
		tasks.GET("", taskAPIService.UpdateTask())
		tasks.GET("/:id", taskAPIService.DeleteTask())
	}
}

func init() {
	setting.Setup()
}

func main() {
	db := postgre.Setup()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	v2 := router.Group("/v2")

	SetupRoutes(db, v2)

	readTimeout := time.Duration(setting.Cfg.Server.ReadTimeout) * time.Second
	writeTimeout := time.Duration(setting.Cfg.Server.WriteTimeout) * time.Second
	endPoint := fmt.Sprintf(":%d", setting.Cfg.Server.Port)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
