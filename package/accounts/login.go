package accounts

import (
	"gather-go/package/accounts/models"
	"gather-go/package/accounts/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) LoginHandle(c *gin.Context) {

	loginBody := LoginBody{}

	if err := c.BindJSON(&loginBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	user := models.User{}

	user.Email = loginBody.Email
	user.Password = loginBody.Password

	token, err := utils.LoginCheck(h.DB, user.Email, user.Password)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}
