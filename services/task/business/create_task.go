package business

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (b *business) CreateTask(ctx context.Context, task *entity.Task) error {
	return b.repo.CreateTask(ctx, task)
}
