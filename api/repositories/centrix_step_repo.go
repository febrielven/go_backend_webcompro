package repositories

import (
	"github.com/febrielven/go_backend_webcompro/api/models"
)

type CentrixStepRepository interface {
	FindAll() ([]*models.CentrixStep, error)
	Save(*models.CentrixStep) (*models.CentrixStep, error)
}
