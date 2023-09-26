package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/modules/user"
)

func InitRouter(userHandler *user.Handler) {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)

	err := r.Run("0.0.0.0:8080")

	if err != nil {
		log.Fatalf("Error init server")
	}
}
