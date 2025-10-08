package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"context"
	"net/http"
	"time"
	"fmt"
	"os"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func main() {
	// 初始化 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),// 先默认本地，后面在 Docker/K8s 改成 redis:6379
	})

	r := gin.Default()

	// 提交任务：POST /submit
	r.POST("/submit", func(c *gin.Context) {
		taskID := fmt.Sprintf("task:%d", time.Now().UnixNano())
		// 模拟任务数据，这里先存字符串 "pending"
		err := rdb.Set(ctx, taskID, "pending", 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 把任务 ID 推进队列
		_ = rdb.LPush(ctx, "task_queue", taskID).Err()

		c.JSON(http.StatusOK, gin.H{"task_id": taskID})
	})

	// 查询任务状态：GET /status/:id
	r.GET("/status/:id", func(c *gin.Context) {
		taskID := c.Param("id")
		val, err := rdb.Get(ctx, taskID).Result()
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": val})
		}
	})

	// 启动服务
	r.Run("0.0.0.0:8080")
}
