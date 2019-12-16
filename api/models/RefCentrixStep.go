package models

import (
	"time"

	"github.com/google/uuid"
)

type CentrixStep struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key;" json:"id"`
	Name      string    `gorm:"column:step;size:20;not null;unique" json:"step"`
	CreatedAt time.Time `gorm:"column:created;default:CURRENT_TIMESTAMP" json:"created"`
}

func (CentrixStep) TableName() string {
	return "ref_centrix_step"
}
