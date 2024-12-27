package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Schedule struct {
	ChannelID   string
	MessageType string
	Time        string
}

var rdb *redis.Client
var ctx context.Context

// Database should contain guild id,
func redis_client_connect_basic() {
	ctx = context.Background()
	envFile, _ := godotenv.Read(".env")
	redisPass, ok := envFile["REDIS_PASS"]
	if !ok {
		log.Fatal("Must set Redis password as env variable: REDIS_PASS")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-17379.c259.us-central1-2.gce.redns.redis-cloud.com:17379",
		Username: "default",
		Password: redisPass,
		DB:       0,
	})

	response, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(response)
	}
}

func setServerSchedule(guildId string, channelId string, timeInterval string, messageType string, timeString string, isOn bool) {
	redis_client_connect_basic()
	key := channelId + timeInterval + messageType
	format := "15:04"
	formattedTime, err := time.Parse(format, timeString)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed time:", formattedTime)
	}
	data := map[string]any{
		"channelId":    channelId,
		"guildId":      guildId,
		"messageType":  messageType,
		"timeInterval": timeInterval,
		"time":         formattedTime,
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, key, serialized, 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	var deserialized []interface{}
	json.Unmarshal([]byte(val), &deserialized)
	fmt.Println("Key: ", key)
	fmt.Println("Value: ", val)
}

func deleteServerSchedule(key string) {
	redis_client_connect_basic()
	deleted, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Key: ", key)
		fmt.Println("Deleted key-value pair ", deleted)
	}
}
