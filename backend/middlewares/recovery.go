package middlewares
import (
	"net/http"
	"hostel-saas/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("Panic recovered", zap.Any("error", err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()
		c.Next()
	}
}
