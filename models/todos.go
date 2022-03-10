package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoStauts string

const (
	TODO_IDLE      TodoStauts = "idle"
	TODO_COMPLETED TodoStauts = "completed"
)

type Todo struct {
	Id        int64          `gorm:"type:integer;primary_key;" json:"id"`
	Uuid      uuid.UUID      `gorm:"type:uuid;" json:"uuid"`
	Title     string         `gorm:"type:string;size:100" json:"title"`
	Content   string         `gorm:"type:string;size:65535" json:"content"`
	Status    TodoStauts     `gorm:"type:string;size:32;default:'idle';check:status IN ('idle', 'completed')" json:"status"`
	CreatedBy int64          `gorm:"foreignKey:Id" json:"created_by"`
	UpdatedBy int64          `gorm:"type:string;foreignKey:Id;" json:"updated_by"`
	DeletedBy int64          `gorm:"type:string;foreignKey:Id;" json:"deleted_by"`
	CreatedAt time.Time      `sql:"DEFAULT:'current_timestamp'" json:"created_at"`
	UpdatedAt time.Time      `sql:"DEFAULT:'current_timestamp'" json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"DEFAULT:'current_timestamp'" json:"deleted_at"`
}
