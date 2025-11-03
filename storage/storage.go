package storage

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"size:128"`
	Sex    string
	Age    int
	Course int
}

type Storage interface {
	Insert(s *Student) error
	Get(id int) (Student, error)
	GetAll() ([]Student, error)
	Update(s *Student) error
	Delete(id int) error
}
