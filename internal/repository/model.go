package repository

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"size:128"`
	Sex    string
	Age    int
	Course int
}
