package handlers

import (
	"net/http"

	"hostel-saas/services"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	service services.AnalyticsService
}

func NewAnalyticsHandler(service services.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{service: service}
}

func (h *AnalyticsHandler) GetDashboardKPIs(c *gin.Context) {
	orgID := getOrgID(c)
	kpis, err := h.service.GetDashboardKPIs(c.Request.Context(), orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get KPIs"})
		return
	}
	c.JSON(http.StatusOK, kpis)
}
