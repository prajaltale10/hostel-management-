package repositories
import (
	"context"
	"hostel-saas/models"
	"gorm.io/gorm"
)
type HostelRepository interface {
	CreateHostel(ctx context.Context, hostel *models.Hostel) error
	GetHostelsByOrgID(ctx context.Context, orgID uint) ([]models.Hostel, error)
	GetHostelHierarchy(ctx context.Context, hostelID uint, orgID uint) (*models.Hostel, error)
	CreateFloor(ctx context.Context, floor *models.Floor) error
	GetFloorByID(ctx context.Context, floorID uint) (*models.Floor, error)
	CreateRoom(ctx context.Context, room *models.Room) error
	GetRoomByID(ctx context.Context, roomID uint) (*models.Room, error)
	CreateBed(ctx context.Context, bed *models.Bed) error
	GetBedByID(ctx context.Context, bedID uint) (*models.Bed, error)
	GetBedByIDAndOrgID(ctx context.Context, bedID uint, orgID uint) (*models.Bed, error)
	UpdateBed(ctx context.Context, bed *models.Bed) error
}
type hostelRepo struct {
	db *gorm.DB
}
func NewHostelRepository(db *gorm.DB) HostelRepository {
	return &hostelRepo{db: db}
}
func (r *hostelRepo) CreateHostel(ctx context.Context, hostel *models.Hostel) error {
	return r.db.WithContext(ctx).Create(hostel).Error
}
func (r *hostelRepo) GetHostelsByOrgID(ctx context.Context, orgID uint) ([]models.Hostel, error) {
	var hostels []models.Hostel
	if err := r.db.WithContext(ctx).Where("org_id = ?", orgID).Find(&hostels).Error; err != nil {
		return nil, err
	}
	return hostels, nil
}
func (r *hostelRepo) GetHostelHierarchy(ctx context.Context, hostelID uint, orgID uint) (*models.Hostel, error) {
	var hostel models.Hostel
	err := r.db.WithContext(ctx).Preload("Floors.Rooms.Beds").Where("id = ? AND org_id = ?", hostelID, orgID).First(&hostel).Error
	if err != nil {
		return nil, err
	}
	return &hostel, nil
}
func (r *hostelRepo) CreateFloor(ctx context.Context, floor *models.Floor) error {
	return r.db.WithContext(ctx).Create(floor).Error
}
func (r *hostelRepo) GetFloorByID(ctx context.Context, floorID uint) (*models.Floor, error) {
	var floor models.Floor
	if err := r.db.WithContext(ctx).First(&floor, floorID).Error; err != nil {
		return nil, err
	}
	return &floor, nil
}
func (r *hostelRepo) CreateRoom(ctx context.Context, room *models.Room) error {
	return r.db.WithContext(ctx).Create(room).Error
}
func (r *hostelRepo) GetRoomByID(ctx context.Context, roomID uint) (*models.Room, error) {
	var room models.Room
	if err := r.db.WithContext(ctx).First(&room, roomID).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
func (r *hostelRepo) CreateBed(ctx context.Context, bed *models.Bed) error {
	return r.db.WithContext(ctx).Create(bed).Error
}
func (r *hostelRepo) GetBedByID(ctx context.Context, bedID uint) (*models.Bed, error) {
	var bed models.Bed
	if err := r.db.WithContext(ctx).First(&bed, bedID).Error; err != nil {
		return nil, err
	}
	return &bed, nil
}
func (r *hostelRepo) GetBedByIDAndOrgID(ctx context.Context, bedID uint, orgID uint) (*models.Bed, error) {
	var bed models.Bed
	err := r.db.WithContext(ctx).
		Joins("JOIN rooms ON rooms.id = beds.room_id").
		Joins("JOIN floors ON floors.id = rooms.floor_id").
		Joins("JOIN hostels ON hostels.id = floors.hostel_id").
		Where("beds.id = ? AND hostels.org_id = ?", bedID, orgID).
		First(&bed).Error
	if err != nil {
		return nil, err
	}
	return &bed, nil
}
func (r *hostelRepo) UpdateBed(ctx context.Context, bed *models.Bed) error {
	return r.db.WithContext(ctx).Save(bed).Error
}
