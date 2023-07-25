package entity

import (
	"net/http"

	"github.com/shinhagunn/micro-clean-architecture-example/helper"
)

var (
	ErrIDInvalid    = helper.NewAPIError(http.StatusBadRequest, "resource.task.id_invalid")
	ErrTaskNotFound = helper.NewAPIError(http.StatusNotFound, "resource.task.not_found")
)
