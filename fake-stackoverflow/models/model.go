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
	Body     string    `json:"body"`
	User     Users     `gorm:"foreignKey:ID" json:"user"`
	Question Questions `gorm:"foreignKey:ID" json:"questions"`
}

type Questions struct {
	gorm.Model
	Tital string `json:"tital"`
	Body  string `json:"body"`
	Tags  []Tags `gorm:"foreignKey:Name" json:"tags"`
	User  Users  `gorm:"foreignKey:ID" json:"users"`
}

type CommentsOnQuestion struct {
	gorm.Model
	Body     string    `json:"body"`
	User     Users     `gorm:"foreignKey:ID" json:"user"`
	Question Questions `gorm:"foreignKey:ID" json:"question"`
}

type CommentsOnAnswer struct {
	gorm.Model
	Body   string  `json:"body"`
	User   Users   `gorm:"foreignKey:ID" json:"user"`
	Answer Answers `gorm:"foreignKey:ID" json:"question"`
}

type Tags struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}
