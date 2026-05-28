package dto
type UpdateOrgRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=Admin Manager Accountant"`
}
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role" binding:"omitempty,oneof=Admin Manager Accountant"`
	IsActive *bool  `json:"is_active"`
}
type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}
