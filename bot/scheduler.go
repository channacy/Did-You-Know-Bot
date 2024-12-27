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

	// rdb.Set(ctx, "foo", "bar", 0)
	// result, err := rdb.Get(ctx, "foo").Result()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result) // >>> bar
}

// key: channel ID, value is server ID, time frequency, at which time, message type (whether it is quote, joke, or advice)
// function that sets the server schedule
// test cases
func setServerSchedule(guildId string, channelId string, timeInterval string, messageType string, isOn bool) {
	redis_client_connect_basic()
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04:05")
	data := map[string]any{
		"guildId":      guildId,
		"messageType":  messageType,
		"timeInterval": timeInterval,
		"time":         formattedTime,
		"isOn":         true,
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, channelId, serialized, 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, channelId).Result()
	if err != nil {
		panic(err)
	}

	var deserialized []interface{}
	json.Unmarshal([]byte(val), &deserialized)
	fmt.Println(deserialized)
}
