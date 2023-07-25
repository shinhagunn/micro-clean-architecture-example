package entity

import "time"

type TaskStatus string

var (
	TaskStatusDoing   = TaskStatus("doing")
	TaskStatusDone    = TaskStatus("done")
	TaskStatusDeleted = TaskStatus("deleted")
)

type Task struct {
	Model

	UserId int64 `gorm:"type:bigint;notnull"`

	Level      int64      `gorm:"type:bigint;notnull"`
	Name       string     `gorm:"type:character varying(30);not null"`
	Status     TaskStatus `gorm:"type:character varying(10);not null;default:doing"`
	DeadlineAt time.Time  `gorm:"type:timestamp;not null"`
}

func (t Task) TableName() string {
	return "tasks"
}
