package business

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

type TaskRepository interface {
	GetTask(ctx context.Context, task *entity.Task, f ...filters.Filter) error
	GetTasks(ctx context.Context, task []entity.Task, f ...filters.Filter) error
	CreateTask(ctx context.Context, task *entity.Task) error
	UpdateTask(ctx context.Context, task *entity.Task, updates *entity.Task) error
	DeleteTask(ctx context.Context, task *entity.Task) error
}

type business struct {
	repo TaskRepository
}

func NewUserBusiness(repo TaskRepository) *business {
	return &business{repo: repo}
}
