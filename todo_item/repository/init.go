package repository

import "gorm.io/gorm"

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repo {
	return &repo{db: db}
}
