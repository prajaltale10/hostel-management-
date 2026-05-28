package redis
import (
	"context"
	"fmt"
	"hostel-saas/config"
	"hostel-saas/pkg/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)
var Client *redis.Client
func InitRedis(cfg *config.Config) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		logger.Log.Error("Failed to connect to Redis", zap.Error(err))
		return err
	}
	logger.Log.Info("Connected to Redis successfully")
	return nil
}
