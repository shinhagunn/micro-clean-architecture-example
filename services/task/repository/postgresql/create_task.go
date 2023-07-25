package postgresql

import (
	"context"

	"github.com/shinhagunn/micro-clean-architecture-example/services/task/entity"
)

func (r *postgresqlRepo) CreateTask(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).Create(&task).Error
}
