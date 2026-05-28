package services
import (
	"context"
	"errors"
	"time"
	"hostel-saas/dto"
	"hostel-saas/models"
	"hostel-saas/repositories"
)
type StudentService interface {
	CreateStudent(ctx context.Context, orgID uint, req dto.CreateStudentRequest) (*dto.StudentResponse, error)
	GetStudent(ctx context.Context, orgID uint, studentID uint) (*dto.StudentResponse, error)
	ListStudents(ctx context.Context, orgID uint) ([]dto.StudentResponse, error)
	UpdateStudent(ctx context.Context, orgID uint, studentID uint, req dto.UpdateStudentRequest) (*dto.StudentResponse, error)
	CheckIn(ctx context.Context, orgID uint, studentID uint, req dto.CheckInRequest) error
	CheckOut(ctx context.Context, orgID uint, studentID uint) error
}
type studentService struct {
	repo       repositories.StudentRepository
	hostelRepo repositories.HostelRepository
}
func NewStudentService(repo repositories.StudentRepository, hostelRepo repositories.HostelRepository) StudentService {
	return &studentService{repo: repo, hostelRepo: hostelRepo}
}
func (s *studentService) CreateStudent(ctx context.Context, orgID uint, req dto.CreateStudentRequest) (*dto.StudentResponse, error) {
	student := &models.Student{OrgID: orgID, Name: req.Name, Email: req.Email, Phone: req.Phone, Address: req.Address, AadharNumber: req.AadharNumber}
	if err := s.repo.CreateStudent(ctx, student); err != nil {
		return nil, err
	}
	return s.mapToResponse(student), nil
}
func (s *studentService) GetStudent(ctx context.Context, orgID uint, studentID uint) (*dto.StudentResponse, error) {
	student, err := s.repo.GetStudentByID(ctx, studentID, orgID)
	if err != nil {
		return nil, err
	}
	return s.mapToResponse(student), nil
}
func (s *studentService) ListStudents(ctx context.Context, orgID uint) ([]dto.StudentResponse, error) {
	students, err := s.repo.ListStudents(ctx, orgID)
	if err != nil {
		return nil, err
	}
	var res []dto.StudentResponse
	for _, st := range students {
		res = append(res, *s.mapToResponse(&st))
	}
	return res, nil
}
func (s *studentService) UpdateStudent(ctx context.Context, orgID uint, studentID uint, req dto.UpdateStudentRequest) (*dto.StudentResponse, error) {
	student, err := s.repo.GetStudentByID(ctx, studentID, orgID)
	if err != nil {
		return nil, err
	}
	if req.Name != "" { student.Name = req.Name }
	if req.Email != "" { student.Email = req.Email }
	if req.Phone != "" { student.Phone = req.Phone }
	if req.Address != "" { student.Address = req.Address }
	if req.AadharNumber != "" { student.AadharNumber = req.AadharNumber }
	if err := s.repo.UpdateStudent(ctx, student); err != nil {
		return nil, err
	}
	return s.mapToResponse(student), nil
}
func (s *studentService) CheckIn(ctx context.Context, orgID uint, studentID uint, req dto.CheckInRequest) error {
	_, err := s.repo.GetStudentByID(ctx, studentID, orgID)
	if err != nil { return errors.New("student not found") }
	return s.repo.WithTransaction(ctx, func(txStudentRepo repositories.StudentRepository, txHostelRepo repositories.HostelRepository) error {
		bed, err := txHostelRepo.GetBedByIDAndOrgID(ctx, req.BedID, orgID)
		if err != nil { return errors.New("bed not found or access denied") }
		if !bed.IsVacant { return errors.New("bed is already occupied") }
		bed.IsVacant = false
		if err := txHostelRepo.UpdateBed(ctx, bed); err != nil { return err }
		allocation := &models.Allocation{StudentID: studentID, BedID: req.BedID, Status: models.StatusActive, CheckIn: req.CheckIn, BaseRent: req.BaseRent}
		if err := txStudentRepo.CreateAllocation(ctx, allocation); err != nil { return err }
		return nil
	})
}
func (s *studentService) CheckOut(ctx context.Context, orgID uint, studentID uint) error {
	_, err := s.repo.GetStudentByID(ctx, studentID, orgID)
	if err != nil { return errors.New("student not found") }
	return s.repo.WithTransaction(ctx, func(txStudentRepo repositories.StudentRepository, txHostelRepo repositories.HostelRepository) error {
		allocation, err := txStudentRepo.GetAllocationByStudentID(ctx, studentID)
		if err != nil { return errors.New("active allocation not found") }
		bed, err := txHostelRepo.GetBedByIDAndOrgID(ctx, allocation.BedID, orgID)
		if err != nil { return errors.New("bed not found or access denied") }
		bed.IsVacant = true
		if err := txHostelRepo.UpdateBed(ctx, bed); err != nil { return err }
		now := time.Now()
		allocation.Status = models.StatusInactive
		allocation.CheckOut = &now
		if err := txStudentRepo.UpdateAllocation(ctx, allocation); err != nil { return err }
		return nil
	})
}
func (s *studentService) mapToResponse(student *models.Student) *dto.StudentResponse {
	return &dto.StudentResponse{ID: student.ID, Name: student.Name, Email: student.Email, Phone: student.Phone, Address: student.Address, AadharNumber: student.AadharNumber}
}
