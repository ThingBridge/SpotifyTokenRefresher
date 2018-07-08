package main

import (
	"fmt"
	"time"

	"./refreshToken"
	"github.com/go-redis/redis"
)

func getAccessToken() {
	accessToken, err := refreshToken.RefreshToken()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err = client.Set("spotify_access_token", accessToken, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println(accessToken)
}

func main() {
	getAccessToken()
	ticker := time.NewTicker(55 * time.Minute)
	go func() {
		for range ticker.C {
			getAccessToken()
		}
	}()
	var response string
	_, _ = fmt.Scanln(&response)
}
