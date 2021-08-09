package entity

import (
	"time"

	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Location *time.Location
}

func New(db *gorm.DB, loc *time.Location) *Database {
	return &Database{
		DB:       db,
		Location: loc,
	}
}

type Owners struct {
	gorm.Model
	todu_id       uint
	role          string `gorm:"default:product owner"`
	name          string
	primary_alias string
}
