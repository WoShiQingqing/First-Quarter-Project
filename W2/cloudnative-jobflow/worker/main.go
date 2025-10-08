package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/redis/go-redis/v9"
)

// 从环境变量获取 Redis 地址（默认 localhost:6379）
func getRedisAddr() string {
    addr := os.Getenv("REDIS_ADDR")
    if addr == "" {
        addr = "localhost:6379"
    }
    return addr
}

func processTask(taskID string) {
    fmt.Printf("Worker: processing %s...\n", taskID)
    time.Sleep(2 * time.Second)
    fmt.Printf("Worker: finished %s\n", taskID)
}

func main() {
    fmt.Println("REDIS_ADDR =", getRedisAddr())

    ctx := context.Background()
    rdb := redis.NewClient(&redis.Options{
        Addr: getRedisAddr(), // ✅ 改这里
    })

    for i := 0; i < 3; i++ {
        go func(id int) {
            for {
                taskID, err := rdb.BRPop(ctx, 0*time.Second, "task_queue").Result()
                if err != nil {
                    log.Println("Worker error:", err)
                    continue
                }
                if len(taskID) > 1 {
                    fmt.Printf("Worker-%d got task: %s\n", id, taskID[1])
                    processTask(taskID[1])
                }
            }
        }(i)
    }

    select {} // 阻塞主线程
}
