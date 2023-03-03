package model

import (
	"time"
)

// WelcomeSortValidValues valid sort values
var WelcomeSortValidValues = map[string]bool{
	"created_at": true,
	"name":       true,
}

// DefaultWelcomeSort a default product sort
var DefaultWelcomeSort = "created_at desc"

// WelcomeRepository :nodoc:
type WelcomeRepository interface {
	Create(welcome *Welcome) error
	FindAll(req *WelcomeQuery) (welcomes *[]Welcome, err error)
	FindByID(id string) (welcome *Welcome, err error)
}

// WelcomeUsecase :nodoc:
type WelcomeUsecase interface {
	Create(welcome *Welcome) (string, error)
	FindAll(req *WelcomeQuery) (products *[]Welcome, err error)
}

// Welcome :nodoc:
type Welcome struct {
	ID        string    `gorm:"type:varchar(128)" json:"id"`
	Message   string    `gorm:"type:varchar(128)" json:"message" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// WelcomeQuery :nodoc:
type WelcomeQuery struct {
	Sort string `json:"sort"`
}
