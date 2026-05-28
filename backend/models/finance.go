package models
import (
	"time"
	"gorm.io/gorm"
)
type RentStatus string
const (
	RentStatusPending RentStatus = "Pending"
	RentStatusPartial RentStatus = "Partial"
	RentStatusPaid    RentStatus = "Paid"
	RentStatusOverdue RentStatus = "Overdue"
)
type Rent struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	StudentID   uint           `gorm:"index;not null" json:"student_id"`
	OrgID       uint           `gorm:"index;not null" json:"org_id"`
	Month       time.Time      `gorm:"not null" json:"month"`
	BaseRent    float64        `gorm:"not null" json:"base_rent"`
	Penalty     float64        `gorm:"default:0" json:"penalty"`
	Discount    float64        `gorm:"default:0" json:"discount"`
	FinalRent   float64        `gorm:"not null" json:"final_rent"`
	PaidAmount  float64        `gorm:"default:0" json:"paid_amount"`
	DueAmount   float64        `gorm:"not null" json:"due_amount"`
	Status      RentStatus     `gorm:"type:varchar(20);default:'Pending'" json:"status"`
	DueDate     time.Time      `gorm:"not null" json:"due_date"`
	Payments    []Payment      `gorm:"foreignKey:RentID" json:"payments,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
type PaymentMethod string
const (
	MethodCash   PaymentMethod = "Cash"
	MethodCard   PaymentMethod = "Card"
	MethodUPI    PaymentMethod = "UPI"
	MethodOnline PaymentMethod = "Online"
)
type Payment struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	RentID        uint           `gorm:"index;not null" json:"rent_id"`
	Amount        float64        `gorm:"not null" json:"amount"`
	PaymentDate   time.Time      `gorm:"not null" json:"payment_date"`
	PaymentMethod PaymentMethod  `gorm:"type:varchar(20);not null" json:"payment_method"`
	TransactionID string         `json:"transaction_id"`
	ReceiptURL    string         `json:"receipt_url"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
