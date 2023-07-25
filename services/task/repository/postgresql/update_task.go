package postgresql

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (r *postgresqlRepo) UpdateTask(ctx context.Context, task *entity.Task, updates *entity.Task) error {
	return r.db.WithContext(ctx).Model(&task).Updates(&updates).Error
}
