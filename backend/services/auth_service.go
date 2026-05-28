package services
import (
	"context"
	"errors"
	"time"
	"hostel-saas/config"
	"hostel-saas/dto"
	"hostel-saas/models"
	"hostel-saas/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
type AuthService interface {
	Signup(ctx context.Context, req dto.SignupRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
}
type authService struct {
	repo repositories.AuthRepository
	cfg  *config.Config
}
func NewAuthService(repo repositories.AuthRepository, cfg *config.Config) AuthService {
	return &authService{repo: repo, cfg: cfg}
}
func (s *authService) Signup(ctx context.Context, req dto.SignupRequest) (*dto.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	org := &models.Organization{
		Name:  req.OrgName,
		Email: req.Email,
	}
	user := &models.User{
		Name:     req.UserName,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     models.RoleOwner,
	}
	if err := s.repo.CreateOrganizationAndOwner(ctx, org, user); err != nil {
		return nil, err
	}
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}
	return s.buildAuthResponse(token, user), nil
}
func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}
	return s.buildAuthResponse(token, user), nil
}
func (s *authService) generateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"org_id":  user.OrgID,
		"role":    string(user.Role),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(s.cfg.JWTSecret))
}
func (s *authService) buildAuthResponse(token string, user *models.User) *dto.AuthResponse {
	return &dto.AuthResponse{
		Token: token,
		User: struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
			Role  string `json:"role"`
			OrgID uint   `json:"org_id"`
		}{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  string(user.Role),
			OrgID: user.OrgID,
		},
	}
}
