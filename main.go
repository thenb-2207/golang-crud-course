package main

import (
  // "net/http"
  "github.com/gin-gonic/gin"
  "example.com/m/v2/models"
  "example.com/m/v2/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDataBase()

  r.GET("/courses", controllers.FindCourses)
  r.POST("/courses", controllers.CreateCourse)
  r.GET("/courses/:id", controllers.FindCourse)
  r.PUT("/courses/:id", controllers.UpdateCourse)
  r.DELETE("/courses/:id", controllers.DeleteCourse)

  r.Run()
}
