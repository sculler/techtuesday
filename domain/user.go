package domain

import (
	apierror "github.com/sculler/techtuesdayapi/error"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	FavoriteColor string `json:"favoriteColor"`
}

type IUserRepository interface {
	Create(user *User) (int, error)
	Delete(id int) (bool, error)
	GetById(id int) (*User, error)
	GetAll() ([]User, error)
	Update(user *User) (bool, error)
}

type IUserService interface {
	Create(user *User) (int, *apierror.ApiError)
	Delete(id int) (bool, *apierror.ApiError)
	GetAll() ([]User, *apierror.ApiError)
	GetById(id int) (*User, *apierror.ApiError)
	Update(user *User) (bool, *apierror.ApiError)
}


