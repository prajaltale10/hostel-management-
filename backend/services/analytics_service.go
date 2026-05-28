package services

import (
	"context"

	"hostel-saas/repositories"
)

type AnalyticsService interface {
	GetDashboardKPIs(ctx context.Context, orgID uint) (map[string]interface{}, error)
}

type analyticsService struct {
	repo repositories.AnalyticsRepository
}

func NewAnalyticsService(repo repositories.AnalyticsRepository) AnalyticsService {
	return &analyticsService{repo: repo}
}

func (s *analyticsService) GetDashboardKPIs(ctx context.Context, orgID uint) (map[string]interface{}, error) {
	return s.repo.GetDashboardKPIs(ctx, orgID)
}
