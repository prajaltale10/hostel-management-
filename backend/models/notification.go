package models
import (
	"time"
	"gorm.io/gorm"
)
type NotificationStatus string
const (
	NotifStatusPending NotificationStatus = "Pending"
	NotifStatusSent    NotificationStatus = "Sent"
	NotifStatusFailed  NotificationStatus = "Failed"
)
type NotificationChannel string
const (
	ChannelEmail    NotificationChannel = "Email"
	ChannelSMS      NotificationChannel = "SMS"
	ChannelWhatsApp NotificationChannel = "WhatsApp"
)
type NotificationLog struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	OrgID     uint                `gorm:"index;not null" json:"org_id"`
	StudentID uint                `gorm:"index" json:"student_id"`
	Channel   NotificationChannel `gorm:"type:varchar(20);not null" json:"channel"`
	Recipient string              `gorm:"not null" json:"recipient"`
	Message   string              `gorm:"type:text;not null" json:"message"`
	Status    NotificationStatus  `gorm:"type:varchar(20);default:'Pending'" json:"status"`
	ErrorMsg  string              `gorm:"type:text" json:"error_msg"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index" json:"-"`
}
