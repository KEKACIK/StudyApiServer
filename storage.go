package main

import (
	"fmt"
	"sync"
)

type Student struct {
	ID     int
	Name   string
	Sex    string
	Age    int
	Course int
}

type Storage interface {
	Insert(s *Student)
	Get(id int) (Student, error)
	GetAll() map[int]Student
	Update(id int, s Student)
	Delete(id int)
	Print()
}
type MemoryStorage struct {
	counter int
	data    map[int]Student
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Student),
		counter: 1,
	}
}

func (st *MemoryStorage) Insert(s *Student) {
	st.Lock()
	s.ID = st.counter
	st.data[s.ID] = *s
	st.counter++
	st.Unlock()
}
func (st *MemoryStorage) Get(id int) (Student, error) {
	student, exists := st.data[id]
	if !exists {
		return Student{}, fmt.Errorf("student id=%d not found", id)
	}
	return student, nil
}
func (st *MemoryStorage) GetAll() map[int]Student {
	return st.data
}

func (st *MemoryStorage) Update(id int, e Student) {

}
func (st *MemoryStorage) Delete(id int) {
	delete(st.data, id)
}
func (st *MemoryStorage) Print() {
	fmt.Println(st.data)
}
