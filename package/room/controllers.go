package room

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/room")

	routes.POST("/create", h.addRoom)
	routes.POST("/create/member", h.addRoomUser)
	routes.GET("rtc/:channelName/:role/:tokenType/:uid/", getRtcToken)
	routes.GET("rtm/:uid/", getRtmToken)
	routes.GET("rte/:channelName/:role/:tokenType/:uid/", getBothTokens)

}
