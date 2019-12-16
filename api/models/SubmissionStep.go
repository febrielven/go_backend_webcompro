package models

import (
	"github.com/google/uuid"
)

type SubmissionStep struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;auto_increment" json:"id"`
	SubmissionStepID uuid.UUID `gorm:"type:uuid;primary_key;not null"; json:"submission_step_id"`
}
