package domain

import (
	apierror "github.com/sculler/techtuesdayapi/error"
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
	Create(techTuesday *TechTuesday) (int, error)
	Delete(id int) (bool, error)
	GetAll() ([]TechTuesday, error)
	GetById(id int) (*TechTuesday, error)
	Update(techTuesday *TechTuesday) (bool, error)
}

type ITechTuesdayService interface {
	Create(techTuesday *TechTuesday) (int, *apierror.ApiError)
	Delete(id int) (bool, *apierror.ApiError)
	GetAll() ([]TechTuesday, *apierror.ApiError)
	GetById(id int) (*TechTuesday, *apierror.ApiError)
	Update(techTuesday *TechTuesday) (bool, *apierror.ApiError)
}