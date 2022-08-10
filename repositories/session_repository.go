package repositories

import (
	"sample_rest_api/models"
	"sample_rest_api/parameters"

	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

type SessionRepository interface {
	GenerateSession(parameters.NewSessionParams) (*models.Session, error)
	FindSessionByID(id uint) (*models.Session, error)
}

func NewSessionRepo(db *gorm.DB) SessionRepository {
	return &sessionRepository{db}
}

func (sr sessionRepository) GenerateSession(params parameters.NewSessionParams) (*models.Session, error) {
	session := models.Session{
		UserID: params.UserID,
	}
	if err := sr.db.Create(&session).Error; err != nil {
		return nil, err

	}
	return &session, nil
}

func (sr sessionRepository) FindSessionByID(id uint) (*models.Session, error) {
	var session models.Session
	if err := sr.db.First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}
