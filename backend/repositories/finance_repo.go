package repositories
import (
	"context"
	"hostel-saas/models"
	"gorm.io/gorm"
)
type FinanceRepository interface {
	CreateRent(ctx context.Context, rent *models.Rent) error
	GetRentByID(ctx context.Context, rentID uint, orgID uint) (*models.Rent, error)
	ListRents(ctx context.Context, orgID uint, status string) ([]models.Rent, error)
	UpdateRent(ctx context.Context, rent *models.Rent) error
	CreatePayment(ctx context.Context, payment *models.Payment) error
	ListPayments(ctx context.Context, orgID uint) ([]models.Payment, error)
	WithTransaction(ctx context.Context, fn func(txRepo FinanceRepository) error) error
}
type financeRepo struct {
	db *gorm.DB
}
func NewFinanceRepository(db *gorm.DB) FinanceRepository {
	return &financeRepo{db: db}
}
func (r *financeRepo) CreateRent(ctx context.Context, rent *models.Rent) error {
	return r.db.WithContext(ctx).Create(rent).Error
}
func (r *financeRepo) GetRentByID(ctx context.Context, rentID uint, orgID uint) (*models.Rent, error) {
	var rent models.Rent
	if err := r.db.WithContext(ctx).Where("id = ? AND org_id = ?", rentID, orgID).First(&rent).Error; err != nil {
		return nil, err
	}
	return &rent, nil
}
func (r *financeRepo) ListRents(ctx context.Context, orgID uint, status string) ([]models.Rent, error) {
	var rents []models.Rent
	query := r.db.WithContext(ctx).Where("org_id = ?", orgID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&rents).Error; err != nil {
		return nil, err
	}
	return rents, nil
}
func (r *financeRepo) UpdateRent(ctx context.Context, rent *models.Rent) error {
	return r.db.WithContext(ctx).Save(rent).Error
}
func (r *financeRepo) CreatePayment(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}
func (r *financeRepo) ListPayments(ctx context.Context, orgID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.WithContext(ctx).Joins("JOIN rents ON rents.id = payments.rent_id").Where("rents.org_id = ?", orgID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
func (r *financeRepo) WithTransaction(ctx context.Context, fn func(txRepo FinanceRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := NewFinanceRepository(tx)
		return fn(txRepo)
	})
}
