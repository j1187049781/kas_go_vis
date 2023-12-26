package models

type Task struct {
	Model
	Time   string  `gorm:"column:time;type:DateTime;not null;comment:time" json:"time"`
	TaskName string  `gorm:"column:task_name;type:VARCHAR(255);not null;comment:task_name" json:"task_name"`
}