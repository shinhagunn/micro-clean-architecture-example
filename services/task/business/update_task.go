package business

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (b *business) UpdateTask(ctx context.Context, task *entity.Task, updates *entity.Task) error {
	return b.repo.UpdateTask(ctx, task, updates)
}
