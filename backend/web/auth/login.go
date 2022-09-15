package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Login
// @Description  Login and get access token
// @Tags         Auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        username formData string true "Username"
// @Param        password formData string true "Password"
// @Router       /auth/login [post]
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
