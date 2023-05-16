package room

import (
	"gather-go/package/accounts/models"
	"gather-go/package/room/room_models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) addRoomUser(c *gin.Context) {
	body := RoomMemberBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var roomMember room_models.RoomMember

	roomMember.UserId = body.UserId

	var user models.User

	if result := h.DB.First(&user, body.UserId); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	roomMember.User = user

	var room room_models.Room

	// if result := h.DB.First(&room, body.RoomUid); result.Error != nil {
	// 	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": result.Error})
	// 	return
	// }
	if err := h.DB.Preload("Owner").Where("id = ?", body.RoomUid).First(&room).Error; err != nil {
		return
	}
	roomMember.Room = room

	if result := h.DB.Create(&roomMember); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Room member could not be created"})
		return
	}

	c.IndentedJSON(http.StatusCreated, &roomMember)
}
