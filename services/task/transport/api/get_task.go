package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

// GET: /tasks/:id
func (api *api) GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			api.ResponseError(c, entity.ErrIDInvalid)
			return
		}

		task, err := api.business.GetTask(
			c.Request.Context(),
			filters.WithID(id),
		)
		if err != nil {
			api.ResponseError(c, entity.ErrTaskNotFound)
			return
		}

		c.JSON(http.StatusOK, task)
	}
}
