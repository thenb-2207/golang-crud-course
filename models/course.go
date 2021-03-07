package models

type Course struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  CourseContents []CourseContent
}
