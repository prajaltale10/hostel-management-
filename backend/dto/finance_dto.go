package dto
import "time"
type GenerateRentRequest struct {
	StudentID uint      `json:"student_id" binding:"required"`
	Month     time.Time `json:"month" binding:"required"`
	BaseRent  float64   `json:"base_rent" binding:"required"`
	Penalty   float64   `json:"penalty"`
	Discount  float64   `json:"discount"`
	DueDate   time.Time `json:"due_date" binding:"required"`
}
type ProcessPaymentRequest struct {
	Amount        float64 `json:"amount" binding:"required"`
	PaymentMethod string  `json:"payment_method" binding:"required,oneof=Cash Card UPI Online"`
	TransactionID string  `json:"transaction_id"`
}
type RentResponse struct {
	ID         uint      `json:"id"`
	StudentID  uint      `json:"student_id"`
	Month      time.Time `json:"month"`
	BaseRent   float64   `json:"base_rent"`
	Penalty    float64   `json:"penalty"`
	Discount   float64   `json:"discount"`
	FinalRent  float64   `json:"final_rent"`
	PaidAmount float64   `json:"paid_amount"`
	DueAmount  float64   `json:"due_amount"`
	Status     string    `json:"status"`
	DueDate    time.Time `json:"due_date"`
}
type PaymentResponse struct {
	ID            uint      `json:"id"`
	RentID        uint      `json:"rent_id"`
	Amount        float64   `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	TransactionID string    `json:"transaction_id"`
	ReceiptURL    string    `json:"receipt_url"`
}
