package room

import (
	"gather-go/package/accounts/models"
)

type RoomBody struct {
	Name       string      `json:"name" binding:"required"`
	OwnerID    uint        `json:"owner_id" binding:"required"`
	Owner      models.User `json:"owner" binding:"required"`
	Start_time string      `json:"start_time" binding:"required"`
	End_time   string      `json:"end_time" binding:"required"`
	Uid        uint        `json:"uid"`
}
