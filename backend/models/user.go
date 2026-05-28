package models
import (
	"time"
	"gorm.io/gorm"
)
type Role string
const (
	RoleOwner     Role = "Owner"
	RoleAdmin     Role = "Admin"
	RoleManager   Role = "Manager"
	RoleAccountant Role = "Accountant"
)
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrgID     uint           `gorm:"index;not null" json:"org_id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Role      Role           `gorm:"type:varchar(20);not null;default:'Manager'" json:"role"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
