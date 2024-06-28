package models

import "gorm.io/gorm"

type ChallengeCategoriesUsersModel struct {
	ID         string
	UserId     string `gorm:"index:idx_catmember"`
	CategoryId string `gorm:"index:idx_catmember"`

	gorm.Model
}
