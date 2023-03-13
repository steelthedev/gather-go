package room

import (
	"gather-go/package/room/room_models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) deleteMember(c *gin.Context) {
	body := RoomMemberBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var room_member room_models.RoomMember

	if err := h.DB.Where("user_id=? AND room_uid=?", body.UserId, body.RoomUid).First(&room_member).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User cannot be found"})
		return
	}
	h.DB.Delete((&room_member))
	c.IndentedJSON(http.StatusOK, gin.H{"messgae": "Member successfully deleted"})
}
