package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/helper"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

// POST: /tasks
func (api *api) CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Payload struct {
			Level int64  `form:"level" json:"level" binding:"required"`
			Name  string `from:"name" json:"name" binding:"required"`
		}

		payload := Payload{}
		if !api.BindAndValid(c, &payload, "resource.task") {
			return
		}

		task := &entity.Task{
			Level: payload.Level,
			Name:  payload.Name,
		}

		if err := api.business.CreateTask(c.Request.Context(), task); err != nil {
			api.ResponseError(c, helper.APIError{Code: http.StatusBadRequest, Err: err})
			return
		}

		api.ResponseData(c, http.StatusCreated, task)
	}
}
