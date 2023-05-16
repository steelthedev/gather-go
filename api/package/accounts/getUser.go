package accounts

import (
	"gather-go/package/accounts/models"
	"gather-go/package/accounts/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) getUser(c *gin.Context) {
	var user models.User

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You need to be logged in"})
		return
	}

	if err := h.DB.First(&user, user_id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User cannot be found"})
		return
	}

	c.IndentedJSON(http.StatusOK, &user)
}
