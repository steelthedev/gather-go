package room

import (
	"gather-go/package/accounts/models"
	"gather-go/package/room/room_models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h handler) addRoom(c *gin.Context) {
	body := RoomBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var room room_models.Room

	room.Name = body.Name
	room.OwnerID = body.OwnerID
	parsed_start, _ := time.Parse("2006-01-02T15:04", string(body.Start_time))
	parsed_end, _ := time.Parse("2006-01-02T15:04", string(body.End_time))
	room.Start_time = parsed_start
	room.End_time = parsed_end
	room.Uid = body.Uid

	var user models.User
	if err := h.DB.First(&user, body.OwnerID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	room.Owner = user

	if result := h.DB.Create(&room); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	c.IndentedJSON(http.StatusCreated, &room)
}
