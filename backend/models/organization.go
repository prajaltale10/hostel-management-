package models
import (
	"time"
	"gorm.io/gorm"
)
type Organization struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Users   []User   `gorm:"foreignKey:OrgID" json:"-"`
	Hostels []Hostel `gorm:"foreignKey:OrgID" json:"-"`
}
