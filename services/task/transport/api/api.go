package api

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/helper"
	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

type TaskBusiness interface {
	GetTask(ctx context.Context, f ...filters.Filter) (*entity.Task, error)
	GetTasks(ctx context.Context, f ...filters.Filter) ([]entity.Task, error)
	CreateTask(ctx context.Context, task *entity.Task) error
	UpdateTask(ctx context.Context, task *entity.Task, updates *entity.Task) error
	DeleteTask(ctx context.Context, task *entity.Task) error
}

type api struct {
	helper.Helper
	business TaskBusiness
}

func NewTaskAPI(business TaskBusiness) *api {
	return &api{business: business}
}
