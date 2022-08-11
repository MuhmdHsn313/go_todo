package repositories

import (
	"github.com/kataras/iris/v12/middleware/jwt"
	"sample_rest_api/models"
	"sample_rest_api/parameters"
	"time"

	"gorm.io/gorm"
)

var (
	secret = []byte("signature_hmac_secret_shared_key")
)

type UserClaims struct {
	jwt.Claims
	UserID uint `json:"user_id"`
}

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
	signer := jwt.NewSigner(jwt.HS256, secret, 10*time.Minute)
	claims := UserClaims{
		UserID: params.UserID,
	}

	token, err := signer.Sign(claims)
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
