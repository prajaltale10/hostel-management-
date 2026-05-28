package handlers
import (
	"net/http"
	"strconv"
	"hostel-saas/dto"
	"hostel-saas/services"
	"github.com/gin-gonic/gin"
)
type FinanceHandler struct {
	service services.FinanceService
}
func NewFinanceHandler(service services.FinanceService) *FinanceHandler {
	return &FinanceHandler{service: service}
}
func (h *FinanceHandler) GenerateRent(c *gin.Context) {
	var req dto.GenerateRentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.GenerateRent(c.Request.Context(), orgID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
func (h *FinanceHandler) ListRents(c *gin.Context) {
	orgID := getOrgID(c)
	status := c.Query("status")
	res, err := h.service.ListRents(c.Request.Context(), orgID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *FinanceHandler) ProcessPayment(c *gin.Context) {
	idStr := c.Param("rent_id")
	rentID, _ := strconv.ParseUint(idStr, 10, 32)
	var req dto.ProcessPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.ProcessPayment(c.Request.Context(), orgID, uint(rentID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
