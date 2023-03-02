package usecase

import (
	"github.com/notblessy/skelago/model"
	"github.com/notblessy/skelago/utils"
	"github.com/sirupsen/logrus"
)

type welcomeUsecase struct {
	welcomeRepo model.WelcomeRepository
}

// NewWelcomeUsecase :nodoc:
func NewWelcomeUsecase(p model.WelcomeRepository) model.WelcomeUsecase {
	return &welcomeUsecase{
		welcomeRepo: p,
	}
}

// Create :nodoc:
func (u *welcomeUsecase) Create(welcome *model.Welcome) (string, error) {
	if welcome.ID == "" {
		welcome.ID = utils.GenerateID()
	}

	err := u.welcomeRepo.Create(welcome)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"welcome": utils.Dump(welcome),
		}).Error(err)

		return "", err
	}

	return welcome.ID, nil
}

// FindAll :nodoc:
func (u *welcomeUsecase) FindAll(req *model.WelcomeQuery) (*[]model.Welcome, error) {
	welcomes, err := u.welcomeRepo.FindAll(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"welcomeQuery": utils.Dump(req),
		}).Error(err)

		return nil, err
	}

	return welcomes, nil
}
