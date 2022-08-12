package repositories

import (
	"github.com/kataras/iris/v12/middleware/jwt"
	"gorm.io/gorm"
	"sample_rest_api/models"
	"sample_rest_api/parameters"
)

type UserClaims struct {
	jwt.Claims
	UserID uint `json:"user_id"`
}

type sessionRepository struct {
	db     *gorm.DB
	signer *jwt.Signer
}

type SessionRepository interface {
	GenerateSession(parameters.NewSessionParams) (*models.Session, error)
	FindSessionByID(id uint) (*models.Session, error)
}

func NewSessionRepo(db *gorm.DB, singer *jwt.Signer) SessionRepository {
	return &sessionRepository{db, singer}
}

func (sr sessionRepository) GenerateSession(params parameters.NewSessionParams) (*models.Session, error) {
	claims := UserClaims{
		UserID: params.UserID,
	}

	token, err := sr.signer.Sign(claims)
	if err != nil {
		return nil, err
	}
	session := models.Session{
		UserID:      params.UserID,
		AccessToken: string(token),
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
