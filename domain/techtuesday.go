package domain

import (
	apierror "github.com/sculler/techtuesday/error"
	"gorm.io/gorm"
	"time"
)

type TechTuesday struct {
	gorm.Model
	Title            string `json:"title"`
	Description      string `json:"description"`
	PresentationDate time.Time `json:"presentationDate"`
	UserId           int `json:"userId"`
	User             User `json:"user"`
}

type ITechTuesdayRepository interface {
	GetAll() ([]TechTuesday, error)
	GetById(id int) (*TechTuesday, error)
	Create(techTuesday *TechTuesday) (int, error)
	Update(techTuesday *TechTuesday) (bool, error)
	Delete(id int) (bool, error)
}

type ITechTuesdayService interface {
	GetAll() ([]TechTuesday, *apierror.ApiError)
	GetById(id int) (*TechTuesday, *apierror.ApiError)
	Create(techTuesday *TechTuesday) (int, *apierror.ApiError)
	Update(techTuesday *TechTuesday) (bool, *apierror.ApiError)
	Delete(id int) (bool, *apierror.ApiError)
}