package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage}
}

func (h *Handler) CreateStudent(c *gin.Context) {
	var student Student
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

	h.storage.Insert(&student)
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
	studentList := h.storage.GetAll()
	c.JSON(http.StatusOK, studentList)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	h.storage.Delete(id)
	c.JSON(http.StatusOK, nil)
}
