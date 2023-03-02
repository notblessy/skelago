package repository

import (
	"fmt"
	"strings"

	"github.com/notblessy/skelago/model"
	"github.com/notblessy/skelago/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type welcomeRepository struct {
	db *gorm.DB
}

// NewWelcomeRepository :nodoc:
func NewWelcomeRepository(d *gorm.DB) model.WelcomeRepository {
	return &welcomeRepository{
		db: d,
	}
}

// Create :nodoc:
func (p *welcomeRepository) Create(welcome *model.Welcome) error {
	logger := logrus.WithFields(logrus.Fields{
		"welcome": utils.Dump(welcome),
	})

	err := p.db.Create(&welcome).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

// FindByID :nodoc:
func (p *welcomeRepository) FindByID(id string) (welcome *model.Welcome, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"id": utils.Dump(id),
	})

	err = p.db.First(&welcome, id).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return welcome, err
}

// FindAll :nodoc:
func (p *welcomeRepository) FindAll(req *model.WelcomeQuery) (welcomes *[]model.Welcome, err error) {
	logger := logrus.WithField("req", req)

	qb := p.db.Model(welcomes)

	if req.Sort != "" {
		sorts := p.sortHandler(req.Sort)
		for _, s := range sorts {
			if s.Field != "" {
				qb.Order(fmt.Sprintf("%s %s", s.Field, s.Type))
			}
		}
	} else {
		qb.Order(model.DefaultWelcomeSort)
	}

	err = qb.Find(&welcomes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err.Error())
		return nil, err
	}

	return welcomes, err
}

func (p *welcomeRepository) sortHandler(sortReq string) (sortResults []utils.Sort) {
	sorts := strings.Split(sortReq, ",")

	if len(sorts) > 0 {
		for _, s := range sorts {
			sortParams := strings.Split(s, "-")

			if len(sortParams) == 2 {
				if _, isSortValidValueExists := model.WelcomeSortValidValues[sortParams[0]]; isSortValidValueExists {
					sortResults = append(sortResults, utils.Sort{
						Field: sortParams[0],
						Type:  strings.ToLower(sortParams[1]),
					})
				}
			}
		}
	}

	return sortResults
}
