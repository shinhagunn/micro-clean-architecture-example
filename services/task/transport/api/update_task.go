package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/helper"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

// PUT: /tasks
func (api *api) UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Payload struct {
			ID    int64  `form:"id" json:"id" binding:"required"`
			Level int64  `form:"level" json:"level"`
			Name  string `from:"name" json:"name"`
		}

		payload := Payload{}
		if ok := api.BindAndValid(c, &payload, "resource.task"); !ok {
			return
		}

		task, err := api.business.GetTask(
			c.Request.Context(),
			// filters.WithFieldEqual("user_id", user.ID),
			filters.WithFieldEqual("id", payload.ID),
		)
		if err != nil {
			api.ResponseError(c, entity.ErrTaskNotFound)
			return
		}

		taskUpdates := &entity.Task{}
		if payload.Level > 0 {
			taskUpdates.Level = payload.Level
		}

		if len(payload.Name) > 0 {
			taskUpdates.Name = payload.Name
		}

		if err := api.business.UpdateTask(c.Request.Context(), task, taskUpdates); err != nil {
			api.ResponseError(c, helper.APIError{Code: http.StatusBadRequest, Err: err})
			return
		}

		api.ResponseData(c, http.StatusOK, task)
	}
}
