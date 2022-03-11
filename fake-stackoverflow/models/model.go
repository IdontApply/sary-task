package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users struct {
	gorm.Model
	Name               string              `gorm:"unique" json:"name"`
	Questions          []Question          `gorm:"foreignKey:UserRefer"`
	Answerss           []Answer            `gorm:"foreignKey:UserRefer"`
	CommentOnQuestions []CommentOnQuestion `gorm:"foreignKey:UserRefer"`
	CommentOnAnswer    []CommentOnAnswer   `gorm:"foreignKey:UserRefer"`
}

type Answer struct {
	gorm.Model
	Body            string            `json:"body" binding:"required"`
	UserRefer       uint              `json:"userrefer" binding:"required"`
	QuestionRefer   uint              `json:"questionrefer" binding:"required"`
	CommentOnAnswer []CommentOnAnswer `gorm:"foreignKey:AnswerRefer"`
}

type Question struct {
	gorm.Model
	Tital              string              `json:"tital" binding:"required"`
	Body               string              `json:"body"`
	Tags               []Tag               `gorm:"many2many:question_tags;"`
	UserRefer          uint                `json:"userrefer" binding:"required"`
	Answers            []Answer            `gorm:"foreignKey:QuestionRefer"`
	CommentOnQuestions []CommentOnQuestion `gorm:"foreignKey:QuestionRefer"`
}

type CommentOnQuestion struct {
	gorm.Model
	Body          string `json:"body" binding:"required"`
	UserRefer     uint   `json:"userrefer" binding:"required"`
	QuestionRefer uint   `json:"questionrefer" binding:"required"`
}

type CommentOnAnswer struct {
	gorm.Model
	Body        string `json:"body"`
	UserRefer   uint   `json:"userrefer" binding:"required"`
	AnswerRefer uint   `json:"answerrefer" binding:"required"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"unique" json:"name" binding:"required"`
}
