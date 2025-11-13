package router

import (
	"StudyApiServer/internal/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	StudyStudentSexMan   = "man"
	StudyStudentSexWoman = "woman"
)

type Router struct {
	gin.Engine
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
		*gin.Default(),
		storage,
		apiToken,
	}

	r.init()

	return &r
}

func (r *Router) init() {
	r.POST("/student", r.CreateStudent)
	r.GET("/student/:id", r.GetStudent)
	r.GET("/student/list", r.GetAllStudent)
	r.PUT("/student/:id", r.UpdateStudent)
	r.DELETE("/student/:id", r.DeleteStudent)
}

func (r *Router) Start() {
	r.Run(":80")
}

func (r *Router) CreateStudent(c *gin.Context) {
	var student repository.Student
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
	if student.Sex != StudyStudentSexMan && student.Sex != StudyStudentSexWoman {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("student 'sex' should be '%s' or '%s'", StudyStudentSexMan, StudyStudentSexWoman),
		})
		return
	}
	if student.Course < 1 || student.Course > 6 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "student 'course' should be between 1 and 6",
		})
		return
	}

	err := r.storage.Insert(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
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
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	student, err := r.storage.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, student)
}

func (r *Router) GetAllStudent(c *gin.Context) {
	studentList, err := r.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, studentList)
}

func (r *Router) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	student, err := r.storage.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	var newStudent repository.Student
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
	if newStudent.Sex == StudyStudentSexMan || newStudent.Sex == StudyStudentSexWoman {
		student.Sex = newStudent.Sex
	}
	if newStudent.Course >= 1 && newStudent.Course <= 6 {
		fmt.Println()
		student.Course = newStudent.Course
	}
	err = r.storage.Update(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, student)
}

func (r *Router) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	err = r.storage.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[int]string{})
}
