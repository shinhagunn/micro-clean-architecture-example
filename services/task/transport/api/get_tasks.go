package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/micro-clean-architecture-example/helper"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
)

// GET: /tasks
func (api *api) GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Payload struct {
			filters.Pagination
			filters.Ordering
		}

		payload := Payload{}
		if !api.BindAndValid(c, &payload, "resource.task") {
			return
		}

		f := []filters.Filter{
			payload.Pagination.GetFilters(),
			payload.Ordering.GetFilters(),
		}

		tasks, err := api.business.GetTasks(c.Request.Context(), f...)
		if err != nil {
			api.ResponseError(c, helper.APIError{Code: http.StatusInternalServerError, Err: err})
			return
		}

		c.JSON(http.StatusOK, tasks)
	}
}
