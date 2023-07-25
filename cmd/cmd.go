package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/composer"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/postgre"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/setting"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start app",
	Run: func(cmd *cobra.Command, args []string) {
		db := postgre.Setup()

		router := gin.New()
		router.Use(gin.Logger(), gin.Recovery())

		v2 := router.Group("/api/v2")

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
	},
}

func SetupRoutes(db *gorm.DB, router *gin.RouterGroup) {
	taskAPIService := composer.ComposerTaskAPIService(db)

	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskAPIService.GetTasks())
		tasks.GET("/:id", taskAPIService.GetTask())
		tasks.POST("", taskAPIService.CreateTask())
		tasks.PUT("", taskAPIService.UpdateTask())
		tasks.DELETE("/:id", taskAPIService.DeleteTask())
	}
}

func init() {
	setting.Setup()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
