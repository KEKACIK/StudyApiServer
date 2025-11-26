package router

import (
	"StudyApiServer/config"
	"StudyApiServer/internal/repository"
	"StudyApiServer/internal/validation"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	storage Storage
	token   string
}

type Storage interface {
	Insert(s *repository.Student) error
	Get(id int) (repository.Student, error)
	GetAll() ([]repository.Student, error)
	Update(s *repository.Student) error
	Delete(id int) error
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewRouter(storage Storage, apiToken string) *Router {
	r := Router{
		gin.Default(),
		storage,
		apiToken,
	}

	r.init()

	return &r
}

func (r *Router) init() {
	r.POST("/student", r.AuthMiddleware, r.CreateStudent)
	r.GET("/student/:id", r.AuthMiddleware, r.GetStudent)
	r.GET("/student/list", r.AuthMiddleware, r.GetAllStudent)
	r.PUT("/student/:id", r.AuthMiddleware, r.UpdateStudent)
	r.DELETE("/student/:id", r.AuthMiddleware, r.DeleteStudent)
}

func (r *Router) Start() {
	r.Run(":80")
}

func (r *Router) CreateStudent(c *gin.Context) {
	var student repository.Student
	if err := c.BindJSON(&student); err != nil {
		fmt.Printf("failed to bind student: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := validation.NameValidation(student.Name); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if err := validation.AgeValidation(student.Age); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if err := validation.SexValidation(student.Sex); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if err := validation.CourseValidation(student.Course); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := r.storage.Insert(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": student.ID,
	})
}

func (r *Router) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	student, err := r.storage.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, student)
}

func (r *Router) GetAllStudent(c *gin.Context) {
	studentList, err := r.storage.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, studentList)
}

func (r *Router) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	student, err := r.storage.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var newStudent repository.Student
	err = c.BindJSON(&newStudent)
	if err != nil {
		fmt.Printf("failed to bind student: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
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
	if newStudent.Sex == config.StudyStudentSexMan || newStudent.Sex == config.StudyStudentSexWoman {
		student.Sex = newStudent.Sex
	}
	if newStudent.Course >= 1 && newStudent.Course <= 6 {
		student.Course = newStudent.Course
	}

	err = r.storage.Update(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, student)
}

func (r *Router) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	err = r.storage.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[int]string{})
}
