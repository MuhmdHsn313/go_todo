package repositories

import (
	"github.com/kataras/iris/v12/x/errors"
	"sample_rest_api/models"
	"sample_rest_api/parameters"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(parameters.NewUserParams) (*models.User, error)
	Login(params parameters.LoginParams) (uint, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uint) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(params parameters.NewUserParams) (*models.User, error) {
	user := models.User{
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Login(params parameters.LoginParams) (uint, error) {
	user, err := ur.FindUserByEmail(params.Email)
	if err != nil {
		return 0, err
	}
	if user.Password == params.Password {
		return user.ID, nil
	}
	return 0, errors.New("incorrect login data")
}
