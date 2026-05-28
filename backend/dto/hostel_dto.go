package dto
type CreateHostelRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
}
type CreateFloorRequest struct {
	Name string `json:"name" binding:"required"`
}
type CreateRoomRequest struct {
	Name string `json:"name" binding:"required"`
}
type CreateBedRequest struct {
	Name string `json:"name" binding:"required"`
}
type HostelResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type BedResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsVacant bool   `json:"is_vacant"`
}
type RoomResponse struct {
	ID   uint          `json:"id"`
	Name string        `json:"name"`
	Beds []BedResponse `json:"beds,omitempty"`
}
type FloorResponse struct {
	ID    uint           `json:"id"`
	Name  string         `json:"name"`
	Rooms []RoomResponse `json:"rooms,omitempty"`
}
type HostelHierarchyResponse struct {
	ID      uint            `json:"id"`
	Name    string          `json:"name"`
	Address string          `json:"address"`
	Floors  []FloorResponse `json:"floors,omitempty"`
}
