package handlers
import (
	"net/http"
	"strconv"
	"hostel-saas/dto"
	"hostel-saas/services"
	"github.com/gin-gonic/gin"
)
type StudentHandler struct {
	service services.StudentService
}
func NewStudentHandler(service services.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orgID := getOrgID(c)
	res, err := h.service.CreateStudent(c.Request.Context(), orgID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
func (h *StudentHandler) GetStudent(c *gin.Context) {
	orgID := getOrgID(c)
	idStr := c.Param("id")
	studentID, _ := strconv.ParseUint(idStr, 10, 32)
	res, err := h.service.GetStudent(c.Request.Context(), orgID, uint(studentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *StudentHandler) ListStudents(c *gin.Context) {
	orgID := getOrgID(c)
	res, err := h.service.ListStudents(c.Request.Context(), orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	orgID := getOrgID(c)
	idStr := c.Param("id")
	studentID, _ := strconv.ParseUint(idStr, 10, 32)
	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.UpdateStudent(c.Request.Context(), orgID, uint(studentID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *StudentHandler) CheckIn(c *gin.Context) {
	orgID := getOrgID(c)
	idStr := c.Param("id")
	studentID, _ := strconv.ParseUint(idStr, 10, 32)
	var req dto.CheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CheckIn(c.Request.Context(), orgID, uint(studentID), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check-in successful"})
}
func (h *StudentHandler) CheckOut(c *gin.Context) {
	orgID := getOrgID(c)
	idStr := c.Param("id")
	studentID, _ := strconv.ParseUint(idStr, 10, 32)
	if err := h.service.CheckOut(c.Request.Context(), orgID, uint(studentID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check-out successful"})
}
