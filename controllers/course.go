package controllers

import (
	"net/http"
  "github.com/gin-gonic/gin"
  "example.com/m/v2/models"
  "gorm.io/gorm"
)

func FindCourses(c *gin.Context) {
  var courses []models.Course
  models.DB.Preload("CourseContents").Find(&courses)

  c.JSON(http.StatusOK, gin.H{"data": courses})
}

func CreateCourse(c *gin.Context) {
  var course models.Course
  if err := c.ShouldBindJSON(&course); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Create(&course)

  c.JSON(http.StatusOK, gin.H{"data": course})
}

func FindCourse(c *gin.Context) {
  var course models.Course

  if err := models.DB.Where("id = ?", c.Param("id")).Preload("CourseContents").First(&course).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": course})
}

func UpdateCourse(c *gin.Context) {
  var course models.Course
  if err := models.DB.Where("id = ?", c.Param("id")).Preload("CourseContents").First(&course).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  if err := c.ShouldBindJSON(&course); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&course)

  c.JSON(http.StatusOK, gin.H{"data": course})
}

func DeleteCourse(c *gin.Context) {
  var course models.Course
  if err := models.DB.Where("id = ?", c.Param("id")).First(&course).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Model(&course).Association("CourseContents").Clear()
  models.DB.Delete(&course)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
