package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/skelago/model"
	"github.com/notblessy/skelago/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type welcomeRepo struct {
	db *gorm.DB
}

var welcomeColumns = []string{"id", "name", "price", "description", "quantity", "created_at", "updated_at"}

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

// TestWelcomeRepo_Create :nodoc:
func TestWelcomeRepo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := initPostgreMock()

	pr := welcomeRepo{
		db: dbMock,
	}

	welcome := &model.Welcome{
		ID:        utils.GenerateID(),
		Message:   "Hello World!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()
		queryResult := sqlmock.NewRows([]string{"id"}).
			AddRow(welcome.ID)
		sqlMock.ExpectQuery("INSERT INTO \"welcomes\"").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(queryResult)
		sqlMock.ExpectCommit()

		err := pr.db.Create(welcome).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()
		queryResult := sqlmock.NewRows([]string{"id"}).
			AddRow(welcome.ID)
		sqlMock.ExpectQuery("INSERT INTO \"welcomes\"").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(queryResult)
		sqlMock.ExpectCommit()

		err := pr.db.Create(welcome).Error
		assert.Error(t, err)
	})

}

// TestWelcomeRepo_FindAll :nodoc:
func TestWelcomeRepo_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := initPostgreMock()

	pr := welcomeRepo{
		db: dbMock,
	}

	welcome := &model.Welcome{
		ID:        utils.GenerateID(),
		Message:   "Hello World!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		rows := sqlmock.NewRows(welcomeColumns).
			AddRow(welcome.ID, welcome.Message, welcome.CreatedAt, welcome.UpdatedAt)

		sqlMock.ExpectQuery("^SELECT .+ FROM \"welcomes\"").
			WillReturnRows(rows)

		err := pr.db.Find(&[]model.Welcome{}).Error
		assert.NoError(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		sqlMock.ExpectQuery("^SELECT .+ FROM \"welcomes\"").
			WillReturnError(gorm.ErrRecordNotFound)

		err := pr.db.Find(&[]model.Welcome{}).Error
		assert.Error(t, err)
	})
}
