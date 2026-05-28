package dto
import "time"
type CreateStudentRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"omitempty,email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	AadharNumber string `json:"aadhar_number"`
}
type UpdateStudentRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email" binding:"omitempty,email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	AadharNumber string `json:"aadhar_number"`
}
type CheckInRequest struct {
	BedID    uint      `json:"bed_id" binding:"required"`
	BaseRent float64   `json:"base_rent" binding:"required"`
	CheckIn  time.Time `json:"check_in" binding:"required"`
}
type StudentResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	AadharNumber string `json:"aadhar_number"`
}
