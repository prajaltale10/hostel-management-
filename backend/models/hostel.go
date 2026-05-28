package models
import (
	"time"
	"gorm.io/gorm"
)
type Hostel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrgID     uint           `gorm:"index;not null" json:"org_id"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `json:"address"`
	Floors    []Floor        `gorm:"foreignKey:HostelID" json:"floors,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Floor struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	HostelID  uint           `gorm:"index;not null" json:"hostel_id"`
	Name      string         `gorm:"not null" json:"name"`
	Rooms     []Room         `gorm:"foreignKey:FloorID" json:"rooms,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Room struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FloorID   uint           `gorm:"index;not null" json:"floor_id"`
	Name      string         `gorm:"not null" json:"name"`
	Beds      []Bed          `gorm:"foreignKey:RoomID" json:"beds,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Bed struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	RoomID    uint           `gorm:"index;not null" json:"room_id"`
	Name      string         `gorm:"not null" json:"name"`
	IsVacant  bool           `gorm:"default:true" json:"is_vacant"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
