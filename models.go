package main

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StuName  string
	StuLevel int32
	StuGPA   float32
}
