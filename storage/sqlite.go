package storage

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteStorage struct {
	DB *gorm.DB
	sync.Mutex
}

func NewSQLiteStorage(databasePath string) (*SQLiteStorage, error) {
	DB, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&Student{})
	return &SQLiteStorage{
		DB: DB,
	}, nil
}

func (st *SQLiteStorage) Insert(s *Student) error {
	st.Lock()
	result := st.DB.Create(&s)
	st.Unlock()
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (st *SQLiteStorage) Get(id int) (Student, error) {
	student := Student{}
	st.Lock()
	st.DB.First(&student, id)
	st.Unlock()
	return student, nil
}

func (st *SQLiteStorage) GetAll() ([]Student, error) {
	students := []Student{}
	st.Lock()
	st.DB.Find(&students)
	st.Unlock()
	return students, nil
}

func (st *SQLiteStorage) Update(s *Student) error {
	st.Lock()
	st.DB.Save(&s)
	st.Unlock()
	return nil
}

func (st *SQLiteStorage) Delete(id int) error {
	student := Student{}
	st.Lock()
	st.DB.First(&student, id)
	st.DB.Delete(&student)
	st.Unlock()
	return nil
}

func (st *SQLiteStorage) Print() {
	fmt.Println(nil)
}
