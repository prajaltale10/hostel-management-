package services
import (
	"context"
	"errors"
	"hostel-saas/dto"
	"hostel-saas/models"
	"hostel-saas/repositories"
	"golang.org/x/crypto/bcrypt"
)
type OrgService interface {
	GetOrganization(ctx context.Context, orgID uint) (*models.Organization, error)
	UpdateOrganization(ctx context.Context, orgID uint, req dto.UpdateOrgRequest) (*models.Organization, error)
	CreateUser(ctx context.Context, orgID uint, req dto.CreateUserRequest) (*dto.UserResponse, error)
	ListUsers(ctx context.Context, orgID uint) ([]dto.UserResponse, error)
	UpdateUser(ctx context.Context, userID uint, orgID uint, req dto.UpdateUserRequest) (*dto.UserResponse, error)
}
type orgService struct {
	repo repositories.OrgRepository
}
func NewOrgService(repo repositories.OrgRepository) OrgService {
	return &orgService{repo: repo}
}
func (s *orgService) GetOrganization(ctx context.Context, orgID uint) (*models.Organization, error) {
	return s.repo.GetOrganizationByID(ctx, orgID)
}
func (s *orgService) UpdateOrganization(ctx context.Context, orgID uint, req dto.UpdateOrgRequest) (*models.Organization, error) {
	org, err := s.repo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		org.Name = req.Name
	}
	if req.Phone != "" {
		org.Phone = req.Phone
	}
	if req.Address != "" {
		org.Address = req.Address
	}
	if err := s.repo.UpdateOrganization(ctx, org); err != nil {
		return nil, err
	}
	return org, nil
}
func (s *orgService) CreateUser(ctx context.Context, orgID uint, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		OrgID:    orgID,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     models.Role(req.Role),
	}
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return s.mapUserToResponse(user), nil
}
func (s *orgService) ListUsers(ctx context.Context, orgID uint) ([]dto.UserResponse, error) {
	users, err := s.repo.GetUsersByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	var responses []dto.UserResponse
	for _, u := range users {
		responses = append(responses, *s.mapUserToResponse(&u))
	}
	return responses, nil
}
func (s *orgService) UpdateUser(ctx context.Context, userID uint, orgID uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, userID, orgID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Role != "" {
		user.Role = models.Role(req.Role)
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}
	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	return s.mapUserToResponse(user), nil
}
func (s *orgService) mapUserToResponse(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     string(user.Role),
		IsActive: user.IsActive,
	}
}
