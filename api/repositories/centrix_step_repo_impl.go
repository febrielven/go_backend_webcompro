package repositories

import (
	"fmt"

	"github.com/febrielven/go_backend_webcompro/api/config"
	"github.com/febrielven/go_backend_webcompro/api/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func NewStepRepository() *stepRepository {
	db := config.GetDB()

	return &stepRepository{
		db: db,
	}
}

type stepRepository struct {
	db *gorm.DB
}

func (server *stepRepository) FindAll() (*[]models.CentrixStep, error) {

	var err error
	step := []models.CentrixStep{}
	//SELECT * FROM users limit 100

	err = server.db.Debug().Model(&models.CentrixStep{}).Find(&step).Error
	if err != nil {
		return &[]models.CentrixStep{}, err
	}

	return &step, err
}

func (server *stepRepository) Save(step *models.CentrixStep) (*models.CentrixStep, error) {
	var err error
	fmt.Printf("data Name: %+v\n", step.Name)

	step.ID = uuid.Must(uuid.NewRandom())

	err = server.db.Debug().Create(&step).Error

	if err != nil {
		return &models.CentrixStep{}, err
	}
	return step, nil
}
