package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/skelago/model"
	"github.com/notblessy/skelago/model/mock"
	"github.com/notblessy/skelago/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgreMock() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return
}

// TestWelcomeUsecase_Create :nodoc:
func TestWelcomeUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	welcomeRepo := mock.NewMockWelcomeRepository(ctrl)

	pr := welcomeUsecase{
		welcomeRepo: welcomeRepo,
	}

	welcome := &model.Welcome{
		ID:      utils.GenerateID(),
		Message: "Hello World!",
	}

	t.Run("Success", func(t *testing.T) {
		welcomeRepo.EXPECT().Create(welcome).
			Times(1).
			Return(nil)

		err := pr.welcomeRepo.Create(welcome)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		welcomeRepo.EXPECT().Create(welcome).
			Times(1).
			Return(errors.New("Internal server error"))

		err := pr.welcomeRepo.Create(welcome)
		assert.Error(t, err)
	})
}

// TestWelcomeUsecase_FindAll :nodoc:
func TestWelcomeUsecase_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	welcomeRepo := mock.NewMockWelcomeRepository(ctrl)

	pr := welcomeUsecase{
		welcomeRepo: welcomeRepo,
	}

	req := &model.WelcomeQuery{
		Sort: "price-desc",
	}

	welcome := &model.Welcome{
		ID:        utils.GenerateID(),
		Message:   "Hello World!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	results := []model.Welcome{}
	results = append(results, *welcome)

	t.Run("Success", func(t *testing.T) {
		welcomeRepo.EXPECT().FindAll(req).
			Times(1).
			Return(&results, nil)

		res, err := pr.welcomeRepo.FindAll(req)
		assert.NoError(t, err)
		assert.Equal(t, &results, res)

	})

	t.Run("Error", func(t *testing.T) {
		welcomeRepo.EXPECT().FindAll(req).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		_, err := pr.welcomeRepo.FindAll(req)
		assert.Error(t, err)
	})
}
