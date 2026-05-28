package services
import (
	"context"
	"errors"
	"hostel-saas/dto"
	"hostel-saas/models"
	"hostel-saas/repositories"
)
type HostelService interface {
	CreateHostel(ctx context.Context, orgID uint, req dto.CreateHostelRequest) (*dto.HostelResponse, error)
	ListHostels(ctx context.Context, orgID uint) ([]dto.HostelResponse, error)
	GetHostelHierarchy(ctx context.Context, orgID uint, hostelID uint) (*dto.HostelHierarchyResponse, error)
	CreateFloor(ctx context.Context, orgID uint, hostelID uint, req dto.CreateFloorRequest) (*models.Floor, error)
	CreateRoom(ctx context.Context, orgID uint, floorID uint, req dto.CreateRoomRequest) (*models.Room, error)
	CreateBed(ctx context.Context, orgID uint, roomID uint, req dto.CreateBedRequest) (*models.Bed, error)
}
type hostelService struct {
	repo repositories.HostelRepository
}
func NewHostelService(repo repositories.HostelRepository) HostelService {
	return &hostelService{repo: repo}
}
func (s *hostelService) CreateHostel(ctx context.Context, orgID uint, req dto.CreateHostelRequest) (*dto.HostelResponse, error) {
	hostel := &models.Hostel{
		OrgID:   orgID,
		Name:    req.Name,
		Address: req.Address,
	}
	if err := s.repo.CreateHostel(ctx, hostel); err != nil {
		return nil, err
	}
	return &dto.HostelResponse{
		ID:      hostel.ID,
		Name:    hostel.Name,
		Address: hostel.Address,
	}, nil
}
func (s *hostelService) ListHostels(ctx context.Context, orgID uint) ([]dto.HostelResponse, error) {
	hostels, err := s.repo.GetHostelsByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	var res []dto.HostelResponse
	for _, h := range hostels {
		res = append(res, dto.HostelResponse{ID: h.ID, Name: h.Name, Address: h.Address})
	}
	return res, nil
}
func (s *hostelService) GetHostelHierarchy(ctx context.Context, orgID uint, hostelID uint) (*dto.HostelHierarchyResponse, error) {
	hostel, err := s.repo.GetHostelHierarchy(ctx, hostelID, orgID)
	if err != nil {
		return nil, errors.New("hostel not found")
	}
	res := &dto.HostelHierarchyResponse{
		ID:      hostel.ID,
		Name:    hostel.Name,
		Address: hostel.Address,
	}
	for _, f := range hostel.Floors {
		floorResp := dto.FloorResponse{ID: f.ID, Name: f.Name}
		for _, r := range f.Rooms {
			roomResp := dto.RoomResponse{ID: r.ID, Name: r.Name}
			for _, b := range r.Beds {
				roomResp.Beds = append(roomResp.Beds, dto.BedResponse{ID: b.ID, Name: b.Name, IsVacant: b.IsVacant})
			}
			floorResp.Rooms = append(floorResp.Rooms, roomResp)
		}
		res.Floors = append(res.Floors, floorResp)
	}
	return res, nil
}
func (s *hostelService) CreateFloor(ctx context.Context, orgID uint, hostelID uint, req dto.CreateFloorRequest) (*models.Floor, error) {
	_, err := s.repo.GetHostelHierarchy(ctx, hostelID, orgID)
	if err != nil {
		return nil, errors.New("hostel not found or not owned by organization")
	}
	floor := &models.Floor{HostelID: hostelID, Name: req.Name}
	if err := s.repo.CreateFloor(ctx, floor); err != nil {
		return nil, err
	}
	return floor, nil
}
func (s *hostelService) CreateRoom(ctx context.Context, orgID uint, floorID uint, req dto.CreateRoomRequest) (*models.Room, error) {
	floor, err := s.repo.GetFloorByID(ctx, floorID)
	if err != nil { return nil, errors.New("floor not found") }
	_, err = s.repo.GetHostelHierarchy(ctx, floor.HostelID, orgID)
	if err != nil { return nil, errors.New("floor's hostel not owned by organization") }

	room := &models.Room{FloorID: floorID, Name: req.Name}
	if err := s.repo.CreateRoom(ctx, room); err != nil {
		return nil, err
	}
	return room, nil
}
func (s *hostelService) CreateBed(ctx context.Context, orgID uint, roomID uint, req dto.CreateBedRequest) (*models.Bed, error) {
	room, err := s.repo.GetRoomByID(ctx, roomID)
	if err != nil { return nil, errors.New("room not found") }
	floor, err := s.repo.GetFloorByID(ctx, room.FloorID)
	if err != nil { return nil, errors.New("floor not found") }
	_, err = s.repo.GetHostelHierarchy(ctx, floor.HostelID, orgID)
	if err != nil { return nil, errors.New("room's hostel not owned by organization") }

	bed := &models.Bed{RoomID: roomID, Name: req.Name, IsVacant: true}
	if err := s.repo.CreateBed(ctx, bed); err != nil {
		return nil, err
	}
	return bed, nil
}
