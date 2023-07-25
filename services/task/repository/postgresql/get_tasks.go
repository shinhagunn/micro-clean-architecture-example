package postgresql

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/pkg/filters"
	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (r *postgresqlRepo) GetTasks(ctx context.Context, task []entity.Task, f ...filters.Filter) error {
	return filters.Apply(r.db.WithContext(ctx).Table(entity.Task{}.TableName()), f).Find(&task).Error
}
