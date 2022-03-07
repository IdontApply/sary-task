package main

import (
	controller "fake-so/controllers"
	m "fake-so/models"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {
	// Create a new user
	r.POST("/user", controller.CreateUser)

	r.POST("/tag", controller.CreateTag)

	r.GET("/questions", controller.GetQuestions)
	r.GET("/question/{q_id}/answers", controller.GetAnswersOfQuestion)
	r.POST("/question", controller.CreateQuestion)
	r.POST("/question/answer", controller.CreateAnswerOfQuestion)

	r.POST("/question/comment", controller.CreateCommentOnQuestion)
	r.POST("/question/answer/comment", controller.CreateCommentOnAnswer)
	r.GET("/question/{q_id}/comments", controller.GetCommentsOnQuestion)
	r.GET("/question/{q_id}/answer/{a_id}/comments", controller.GetCommentsOnAnswer)
}

// @title           Fake StackOverflow
// @version         1.0

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	r := gin.Default()
	InitializeRouter(r)
	m.ConnectDatabase()

	r.Run()
}
