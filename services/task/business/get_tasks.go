package business

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (b *business) GetTasks(ctx context.Context, f ...filters.Filter) ([]entity.Task, error) {
	tasks := make([]entity.Task, 0)

	if err := b.repo.GetTasks(ctx, tasks, f...); err != nil {
		return nil, err
	}

	return tasks, nil
}
