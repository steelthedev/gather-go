package room

import (
	"gather-go/package/room/room_models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) getMember(c *gin.Context) {

	body := RoomMemberBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var room_member room_models.RoomMember

	if err := h.DB.Where("userId=?, roomUid=?", body.UserId, body.RoomUid).First(&room_member); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, &room_member)
}
