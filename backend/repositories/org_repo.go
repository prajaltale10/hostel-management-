package repositories
import (
	"context"
	"hostel-saas/models"
	"gorm.io/gorm"
)
type OrgRepository interface {
	GetOrganizationByID(ctx context.Context, id uint) (*models.Organization, error)
	UpdateOrganization(ctx context.Context, org *models.Organization) error
	CreateUser(ctx context.Context, user *models.User) error
	GetUsersByOrgID(ctx context.Context, orgID uint) ([]models.User, error)
	GetUserByID(ctx context.Context, id uint, orgID uint) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}
type orgRepo struct {
	db *gorm.DB
}
func NewOrgRepository(db *gorm.DB) OrgRepository {
	return &orgRepo{db: db}
}
func (r *orgRepo) GetOrganizationByID(ctx context.Context, id uint) (*models.Organization, error) {
	var org models.Organization
	if err := r.db.WithContext(ctx).First(&org, id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}
func (r *orgRepo) UpdateOrganization(ctx context.Context, org *models.Organization) error {
	return r.db.WithContext(ctx).Save(org).Error
}
func (r *orgRepo) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
func (r *orgRepo) GetUsersByOrgID(ctx context.Context, orgID uint) ([]models.User, error) {
	var users []models.User
	if err := r.db.WithContext(ctx).Where("org_id = ?", orgID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *orgRepo) GetUserByID(ctx context.Context, id uint, orgID uint) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("id = ? AND org_id = ?", id, orgID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *orgRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}
