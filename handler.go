package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
		fmt.Printf("Failed to bind student: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if student.Name == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Student 'Name' is empty",
		})
		return
	}
	if student.Age == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Student 'Age' is empty",
		})
		return
	}
	if student.Sex != "Man" && student.Sex != "Woman" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Student 'Sex' should be 'Man' or 'Woman'",
		})
		return
	}
	if student.Course < 1 || student.Course > 6 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Student 'Course' should be between 1 and 6",
		})
		return
	}

	h.storage.Insert(&student)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": student.ID,
	})
}
