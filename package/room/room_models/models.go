package room_models

import (
	"gather-go/package/accounts/models"
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID         uint        `gorm:"primary_key;index:idx_name,unique"`
	Name       string      `json:"name"`
	OwnerID    uint        `json:"owner_id"`
	Owner      models.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Start_time time.Time   `json:"start_time"`
	End_time   time.Time   `json:"end_time"`
	Uid        uint        `json:"uid"`
}

type RoomMember struct {
	gorm.Model

	UserId  uint        `json:"user_id"`
	User    models.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomUid uint        `json:"room_uid"`
	Room    Room        `json:"room" gorm:"foreignKey:RoomUid;references:ID"`
}
