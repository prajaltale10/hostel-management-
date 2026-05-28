package repositories
import (
	"context"
	"hostel-saas/models"
	"gorm.io/gorm"
)
type StudentRepository interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	GetStudentByID(ctx context.Context, studentID uint, orgID uint) (*models.Student, error)
	ListStudents(ctx context.Context, orgID uint) ([]models.Student, error)
	UpdateStudent(ctx context.Context, student *models.Student) error
	CreateAllocation(ctx context.Context, allocation *models.Allocation) error
	GetAllocationByStudentID(ctx context.Context, studentID uint) (*models.Allocation, error)
	UpdateAllocation(ctx context.Context, allocation *models.Allocation) error
	WithTransaction(ctx context.Context, fn func(txRepo StudentRepository, hostelRepo HostelRepository) error) error
}
type studentRepo struct {
	db *gorm.DB
}
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepo{db: db}
}
func (r *studentRepo) CreateStudent(ctx context.Context, student *models.Student) error {
	return r.db.WithContext(ctx).Create(student).Error
}
func (r *studentRepo) GetStudentByID(ctx context.Context, studentID uint, orgID uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.WithContext(ctx).Preload("Allocations").Where("id = ? AND org_id = ?", studentID, orgID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
func (r *studentRepo) ListStudents(ctx context.Context, orgID uint) ([]models.Student, error) {
	var students []models.Student
	if err := r.db.WithContext(ctx).Where("org_id = ?", orgID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
func (r *studentRepo) UpdateStudent(ctx context.Context, student *models.Student) error {
	return r.db.WithContext(ctx).Save(student).Error
}
func (r *studentRepo) CreateAllocation(ctx context.Context, allocation *models.Allocation) error {
	return r.db.WithContext(ctx).Create(allocation).Error
}
func (r *studentRepo) GetAllocationByStudentID(ctx context.Context, studentID uint) (*models.Allocation, error) {
	var allocation models.Allocation
	if err := r.db.WithContext(ctx).Where("student_id = ? AND status = ?", studentID, models.StatusActive).First(&allocation).Error; err != nil {
		return nil, err
	}
	return &allocation, nil
}
func (r *studentRepo) UpdateAllocation(ctx context.Context, allocation *models.Allocation) error {
	return r.db.WithContext(ctx).Save(allocation).Error
}
func (r *studentRepo) WithTransaction(ctx context.Context, fn func(txRepo StudentRepository, hostelRepo HostelRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txStudentRepo := NewStudentRepository(tx)
		txHostelRepo := NewHostelRepository(tx)
		return fn(txStudentRepo, txHostelRepo)
	})
}
