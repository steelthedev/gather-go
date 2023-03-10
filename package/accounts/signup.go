package accounts

import (
	"gather-go/package/accounts/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) SignUp(c *gin.Context) {

	body := createUserBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var user models.User

	user.Username = body.Username
	user.Email = body.Email
	user.Full_name = body.Full_name
	hashedpwd, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	user.Password = string(hashedpwd)
	user.Phone_number = body.Phone_number

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, &user)
}
