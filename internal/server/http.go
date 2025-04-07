package server

import (
	"net/http"

	"github.com/alirezaghasemi/hexagonal_arch/internal/entity"
	"github.com/alirezaghasemi/hexagonal_arch/internal/usecase/general"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userUC general.UserUseCase) {
	r.POST("/api/v1/users", func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		created, err := userUC.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, created)
	})

	r.GET("/api/v1/users/:email", func(c *gin.Context) {
		email := c.Param("email")

		user, err := userUC.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})
}
