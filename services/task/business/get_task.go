package business

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (b *business) GetTask(ctx context.Context, f ...filters.Filter) (*entity.Task, error) {
	var task *entity.Task

	if err := b.repo.GetTask(ctx, task, f...); err != nil {
		return nil, err
	}

	return task, nil
}
