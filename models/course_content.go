package models

type CourseContent struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Description string `json:"description"`
  CourseID uint
}
