package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestCreateSessionRepository(t *testing.T) {
	_, _, err := GetDB()
	if err != nil {
		t.Error(err)
		return
	}
}

const UserId = 1

// func TestCreateNewSession(t *testing.T) {
// 	db, mock, err := GetDB()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	repo := NewSessionRepo(db)
// 	mock.ExpectPrepare("INSERT INTO SESSION (USER_ID,ACCESS_TOKEN) VALUES (1,'abcd') RETURN *")
// 	session, err := repo.GenerateSession(parameters.NewSessionParams{UserID: UserId})
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	t.Log(session.AccessToken)

// }
