package config

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Conn() {
	// Create a new Redis client.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // Set if your Redis server requires authentication
		DB:       0,                // Use the default database
	})

	// Set a key-value pair in Redis.
	err := rdb.Set(context.Background(), "name", "John Doe", 0).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}

	// Get the value associated with the key "name" from Redis.
	name, err := rdb.Get(context.Background(), "name").Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}

	fmt.Println("Name:", name)

	// Set a key-value pair with expiration (in this case, 1 minute).
	err = rdb.Set(context.Background(), "city", "New York", time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value with expiration:", err)
		return
	}

	// Sleep for 2 seconds to allow the key to expire.
	time.Sleep(2 * time.Second)

	// Attempt to get the value of the key "city" after the expiration time.
	city, err := rdb.Get(context.Background(), "city").Result()
	if err == redis.Nil {
		fmt.Println("Key 'city' has expired.")
	} else if err != nil {
		fmt.Println("Error getting value:", err)
		return
	} else {
		fmt.Println("City:", city)
	}
}
