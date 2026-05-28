package repositories
import (
	"context"
	"hostel-saas/models"
	"gorm.io/gorm"
)
type AuthRepository interface {
	CreateOrganizationAndOwner(ctx context.Context, org *models.Organization, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
type authRepo struct {
	db *gorm.DB
}
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepo{db: db}
}
func (r *authRepo) CreateOrganizationAndOwner(ctx context.Context, org *models.Organization, user *models.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(org).Error; err != nil {
			return err
		}
		user.OrgID = org.ID
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
}
func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
