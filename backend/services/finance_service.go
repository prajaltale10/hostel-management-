package services
import (
	"context"
	"errors"
	"time"
	"hostel-saas/dto"
	"hostel-saas/models"
	"hostel-saas/repositories"
)
type FinanceService interface {
	GenerateRent(ctx context.Context, orgID uint, req dto.GenerateRentRequest) (*dto.RentResponse, error)
	ListRents(ctx context.Context, orgID uint, status string) ([]dto.RentResponse, error)
	ProcessPayment(ctx context.Context, orgID uint, rentID uint, req dto.ProcessPaymentRequest) (*dto.PaymentResponse, error)
}
type financeService struct {
	repo repositories.FinanceRepository
}
func NewFinanceService(repo repositories.FinanceRepository) FinanceService {
	return &financeService{repo: repo}
}
func (s *financeService) GenerateRent(ctx context.Context, orgID uint, req dto.GenerateRentRequest) (*dto.RentResponse, error) {
	finalRent := req.BaseRent + req.Penalty - req.Discount
	rent := &models.Rent{
		StudentID: req.StudentID,
		OrgID:     orgID,
		Month:     req.Month,
		BaseRent:  req.BaseRent,
		Penalty:   req.Penalty,
		Discount:  req.Discount,
		FinalRent: finalRent,
		DueAmount: finalRent,
		Status:    models.RentStatusPending,
		DueDate:   req.DueDate,
	}
	if err := s.repo.CreateRent(ctx, rent); err != nil {
		return nil, err
	}
	return s.mapRentToResponse(rent), nil
}
func (s *financeService) ListRents(ctx context.Context, orgID uint, status string) ([]dto.RentResponse, error) {
	rents, err := s.repo.ListRents(ctx, orgID, status)
	if err != nil {
		return nil, err
	}
	var res []dto.RentResponse
	for _, r := range rents {
		res = append(res, *s.mapRentToResponse(&r))
	}
	return res, nil
}
func (s *financeService) ProcessPayment(ctx context.Context, orgID uint, rentID uint, req dto.ProcessPaymentRequest) (*dto.PaymentResponse, error) {
	var paymentResp *dto.PaymentResponse
	err := s.repo.WithTransaction(ctx, func(txRepo repositories.FinanceRepository) error {
		rent, err := txRepo.GetRentByID(ctx, rentID, orgID)
		if err != nil {
			return errors.New("rent record not found")
		}
		if rent.Status == models.RentStatusPaid {
			return errors.New("rent is already fully paid")
		}
		if req.Amount > rent.DueAmount {
			return errors.New("payment amount exceeds due amount")
		}
		payment := &models.Payment{
			RentID:        rent.ID,
			Amount:        req.Amount,
			PaymentDate:   time.Now(),
			PaymentMethod: models.PaymentMethod(req.PaymentMethod),
			TransactionID: req.TransactionID,
		}
		if err := txRepo.CreatePayment(ctx, payment); err != nil {
			return err
		}
		rent.PaidAmount += req.Amount
		rent.DueAmount = rent.FinalRent - rent.PaidAmount
		if rent.DueAmount == 0 {
			rent.Status = models.RentStatusPaid
		} else {
			rent.Status = models.RentStatusPartial
		}
		if err := txRepo.UpdateRent(ctx, rent); err != nil {
			return err
		}
		paymentResp = &dto.PaymentResponse{
			ID:            payment.ID,
			RentID:        payment.RentID,
			Amount:        payment.Amount,
			PaymentDate:   payment.PaymentDate,
			PaymentMethod: string(payment.PaymentMethod),
			TransactionID: payment.TransactionID,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paymentResp, nil
}
func (s *financeService) mapRentToResponse(rent *models.Rent) *dto.RentResponse {
	return &dto.RentResponse{
		ID:         rent.ID,
		StudentID:  rent.StudentID,
		Month:      rent.Month,
		BaseRent:   rent.BaseRent,
		Penalty:    rent.Penalty,
		Discount:   rent.Discount,
		FinalRent:  rent.FinalRent,
		PaidAmount: rent.PaidAmount,
		DueAmount:  rent.DueAmount,
		Status:     string(rent.Status),
		DueDate:    rent.DueDate,
	}
}
