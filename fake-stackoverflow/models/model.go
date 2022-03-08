package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}

type Answers struct {
	gorm.Model
	Body     string    `json:"body" binding:"required"`
	User     Users     `gorm:"foreignKey:ID" json:"user" binding:"required"`
	Question Questions `gorm:"foreignKey:ID" json:"questions" binding:"required"`
}

type Questions struct {
	gorm.Model
	Tital string `json:"tital" binding:"required"`
	Body  string `json:"body"`
	Tags  []Tags `gorm:"foreignKey:Name" json:"tags"`
	User  Users  `gorm:"foreignKey:ID" json:"users" binding:"required"`
}

type CommentsOnQuestion struct {
	gorm.Model
	Body     string    `json:"body" binding:"required"`
	User     Users     `gorm:"foreignKey:ID" json:"user" binding:"required"`
	Question Questions `gorm:"foreignKey:ID" json:"question" binding:"required"`
}

type CommentsOnAnswer struct {
	gorm.Model
	Body   string  `json:"body"`
	User   Users   `gorm:"foreignKey:ID" json:"user" binding:"required"`
	Answer Answers `gorm:"foreignKey:ID" json:"question" binding:"required"`
}

type Tags struct {
	gorm.Model
	Name string `gorm:"unique" json:"name" binding:"required"`
}
