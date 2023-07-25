package postgresql

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (r *postgresqlRepo) DeleteTask(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).Model(task).Update("state", entity.TaskStatusDeleted).Error
}
