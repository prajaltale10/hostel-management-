package handlers
import (
	"net/http"
	"strconv"
	"hostel-saas/dto"
	"hostel-saas/services"
	"github.com/gin-gonic/gin"
)
type OrgHandler struct {
	service services.OrgService
}
func NewOrgHandler(service services.OrgService) *OrgHandler {
	return &OrgHandler{service: service}
}
func getOrgID(c *gin.Context) uint {
	val, _ := c.Get("orgID")
	switch v := val.(type) {
	case float64:
		return uint(v)
	case uint:
		return v
	default:
		return 0
	}
}
func (h *OrgHandler) GetOrganization(c *gin.Context) {
	orgID := getOrgID(c)
	org, err := h.service.GetOrganization(c.Request.Context(), orgID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	c.JSON(http.StatusOK, org)
}
func (h *OrgHandler) UpdateOrganization(c *gin.Context) {
	var req dto.UpdateOrgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	org, err := h.service.UpdateOrganization(c.Request.Context(), orgID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, org)
}
func (h *OrgHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	user, err := h.service.CreateUser(c.Request.Context(), orgID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
func (h *OrgHandler) ListUsers(c *gin.Context) {
	orgID := getOrgID(c)
	users, err := h.service.ListUsers(c.Request.Context(), orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (h *OrgHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	user, err := h.service.UpdateUser(c.Request.Context(), uint(userID), orgID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
