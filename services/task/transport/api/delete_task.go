package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/helper"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

// DELETE: /tasks
func (api *api) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Payload struct {
			ID int64 `form:"id" json:"id" binding:"required"`
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

		if err := api.business.DeleteTask(c.Request.Context(), task); err != nil {
			api.ResponseError(c, helper.APIError{Code: http.StatusBadRequest, Err: err})
			return
		}

		c.JSON(http.StatusOK, 200)
	}
}
