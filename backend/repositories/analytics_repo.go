package repositories

import (
	"context"
	"hostel-saas/models"

	"gorm.io/gorm"
)

type AnalyticsRepository interface {
	GetDashboardKPIs(ctx context.Context, orgID uint) (map[string]interface{}, error)
}

type analyticsRepo struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) AnalyticsRepository {
	return &analyticsRepo{db: db}
}

func (r *analyticsRepo) GetDashboardKPIs(ctx context.Context, orgID uint) (map[string]interface{}, error) {
	var totalStudents int64
	r.db.WithContext(ctx).Model(&models.Student{}).Where("org_id = ?", orgID).Count(&totalStudents)

	var occupiedBeds int64
	r.db.WithContext(ctx).Model(&models.Bed{}).
		Joins("JOIN rooms ON rooms.id = beds.room_id").
		Joins("JOIN floors ON floors.id = rooms.floor_id").
		Joins("JOIN hostels ON hostels.id = floors.hostel_id").
		Where("hostels.org_id = ? AND beds.is_vacant = ?", orgID, false).Count(&occupiedBeds)

	var vacantBeds int64
	r.db.WithContext(ctx).Model(&models.Bed{}).
		Joins("JOIN rooms ON rooms.id = beds.room_id").
		Joins("JOIN floors ON floors.id = rooms.floor_id").
		Joins("JOIN hostels ON hostels.id = floors.hostel_id").
		Where("hostels.org_id = ? AND beds.is_vacant = ?", orgID, true).Count(&vacantBeds)

	var totalRevenue float64
	r.db.WithContext(ctx).Model(&models.Rent{}).Where("org_id = ?", orgID).Select("COALESCE(SUM(paid_amount), 0)").Scan(&totalRevenue)

	var pendingDues float64
	r.db.WithContext(ctx).Model(&models.Rent{}).Where("org_id = ?", orgID).Select("COALESCE(SUM(due_amount), 0)").Scan(&pendingDues)

	kpis := map[string]interface{}{
		"total_students": totalStudents,
		"occupied_beds":  occupiedBeds,
		"vacant_beds":    vacantBeds,
		"total_revenue":  totalRevenue,
		"pending_dues":   pendingDues,
	}

	return kpis, nil
}
