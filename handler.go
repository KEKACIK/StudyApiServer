package main

import (
	"StudyApiServer/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
	return &Handler{storage}
}

func (h *Handler) CreateStudent(c *gin.Context) {
	var student storage.Student
	if err := c.BindJSON(&student); err != nil {
		fmt.Printf("failed to bind student: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if student.Name == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "student 'name' is empty",
		})
		return
	}
	if student.Age == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "student 'age' is empty",
		})
		return
	}
	if student.Sex != "Man" && student.Sex != "Woman" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "student 'sex' should be 'Man' or 'Woman'",
		})
		return
	}
	if student.Course < 1 || student.Course > 6 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "student 'course' should be between 1 and 6",
		})
		return
	}

	err := h.storage.Insert(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": student.ID,
	})
}

func (h *Handler) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	student, err := h.storage.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, student)
}

func (h *Handler) GetAllStudent(c *gin.Context) {
	studentList, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, studentList)
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	student, err := h.storage.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	var newStudent storage.Student
	if err := c.BindJSON(&newStudent); err != nil {
		fmt.Printf("failed to bind student: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if newStudent.Name != "" {
		student.Name = newStudent.Name
	}
	if newStudent.Age != 0 {
		student.Age = newStudent.Age
	}
	if newStudent.Sex == "Man" || newStudent.Sex == "Woman" {
		student.Sex = newStudent.Sex
	}
	if newStudent.Course > 1 && newStudent.Course < 6 {
		fmt.Println()
		student.Course = newStudent.Course
	}
	err = h.storage.Update(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, student)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	err = h.storage.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[int]string{})
}
