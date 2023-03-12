package room_models

import (
	"gather-go/package/accounts/models"
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model

	Name       string      `json:"username"`
	OwnerID    uint        `json:"owner_id"`
	Owner      models.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Start_time time.Time   `json:"start_time"`
	End_time   time.Time   `json:"end_time"`
	Uid        uint        `json:"uid" gorm:"primary_key"`
}

// h.DB.Model(&Child{}).AddForeignKey("parent_id", "parents(id)", "CASCADE", "CASCADE")
