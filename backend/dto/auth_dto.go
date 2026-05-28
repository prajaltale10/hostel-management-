package dto
type SignupRequest struct {
	OrgName  string `json:"org_name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
		OrgID uint   `json:"org_id"`
	} `json:"user"`
}
