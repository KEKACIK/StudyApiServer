package main

import (
	"StudyApiServer/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	sqliteStorage, err := storage.NewSQLiteStorage("database.db")
	if err != nil {
		panic(err)
	}

	handler := NewHandler(sqliteStorage)
	router := gin.Default()

	router.POST("/student", handler.CreateStudent)
	router.GET("/student/:id", handler.GetStudent)
	router.GET("/student/list", handler.GetAllStudent)
	router.PUT("/student/:id", handler.UpdateStudent)
	router.DELETE("/student/:id", handler.DeleteStudent)

	router.Run(":80")
}
