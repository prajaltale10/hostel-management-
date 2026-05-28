package handlers
import (
	"net/http"
	"strconv"
	"hostel-saas/dto"
	"hostel-saas/services"
	"github.com/gin-gonic/gin"
)
type HostelHandler struct {
	service services.HostelService
}
func NewHostelHandler(service services.HostelService) *HostelHandler {
	return &HostelHandler{service: service}
}
func (h *HostelHandler) CreateHostel(c *gin.Context) {
	var req dto.CreateHostelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.CreateHostel(c.Request.Context(), orgID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
func (h *HostelHandler) ListHostels(c *gin.Context) {
	orgID := getOrgID(c)
	res, err := h.service.ListHostels(c.Request.Context(), orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *HostelHandler) GetHostelHierarchy(c *gin.Context) {
	orgID := getOrgID(c)
	idStr := c.Param("id")
	hostelID, _ := strconv.ParseUint(idStr, 10, 32)
	res, err := h.service.GetHostelHierarchy(c.Request.Context(), orgID, uint(hostelID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *HostelHandler) CreateFloor(c *gin.Context) {
	idStr := c.Param("id")
	hostelID, _ := strconv.ParseUint(idStr, 10, 32)
	var req dto.CreateFloorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.CreateFloor(c.Request.Context(), orgID, uint(hostelID), req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
func (h *HostelHandler) CreateRoom(c *gin.Context) {
	idStr := c.Param("floor_id")
	floorID, _ := strconv.ParseUint(idStr, 10, 32)
	var req dto.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.CreateRoom(c.Request.Context(), orgID, uint(floorID), req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
func (h *HostelHandler) CreateBed(c *gin.Context) {
	idStr := c.Param("room_id")
	roomID, _ := strconv.ParseUint(idStr, 10, 32)
	var req dto.CreateBedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.CreateBed(c.Request.Context(), orgID, uint(roomID), req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
