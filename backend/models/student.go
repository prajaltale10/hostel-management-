package models
import (
	"time"
	"gorm.io/gorm"
)
type Student struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	OrgID        uint           `gorm:"index;not null" json:"org_id"`
	Name         string         `gorm:"not null" json:"name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Address      string         `json:"address"`
	AadharNumber string         `json:"aadhar_number"`
	Allocations  []Allocation   `gorm:"foreignKey:StudentID" json:"allocations,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
type AllocationStatus string
const (
	StatusActive   AllocationStatus = "Active"
	StatusInactive AllocationStatus = "Inactive"
)
type Allocation struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	StudentID uint             `gorm:"index;not null" json:"student_id"`
	BedID     uint             `gorm:"index;not null" json:"bed_id"`
	Status    AllocationStatus `gorm:"type:varchar(20);default:'Active'" json:"status"`
	CheckIn   time.Time        `gorm:"not null" json:"check_in"`
	CheckOut  *time.Time       `json:"check_out"`
	BaseRent  float64          `gorm:"not null" json:"base_rent"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}
