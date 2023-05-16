package room

import (
	"gather-go/package/accounts/models"
	"gather-go/package/room/room_models"
)

type RoomBody struct {
	Name       string      `json:"name" binding:"required"`
	OwnerID    uint        `json:"owner_id" binding:"required"`
	Owner      models.User `json:"owner" binding:"required"`
	Start_time string      `json:"start_time" binding:"required"`
	End_time   string      `json:"end_time" binding:"required"`
	Uid        uint        `json:"uid"`
}

type RoomMemberBody struct {
	UserId  uint             `json:"user_id"`
	User    models.User      `json:"user"`
	RoomUid uint             `json:"room_uid"`
	Room    room_models.Room `json:"room_name"`
}
