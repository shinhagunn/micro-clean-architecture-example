package composer

import (
	"github.com/gin-gonic/gin"
	taskBusiness "github.com/shinhagunn/micro-clean-architecture-example/services/task/business"
	taskRepo "github.com/shinhagunn/micro-clean-architecture-example/services/task/repository/postgresql"
	taskAPI "github.com/shinhagunn/micro-clean-architecture-example/services/task/transport/api"
	"gorm.io/gorm"
)

type TaskComposer interface {
	GetTasks() gin.HandlerFunc
	GetTask() gin.HandlerFunc
	CreateTask() gin.HandlerFunc
	UpdateTask() gin.HandlerFunc
	DeleteTask() gin.HandlerFunc
}

func ComposerTaskAPIService(db *gorm.DB) TaskComposer {
	repo := taskRepo.NewPostgresqlRepo(db)
	business := taskBusiness.NewUserBusiness(repo)
	api := taskAPI.NewTaskAPI(business)

	return api
}
