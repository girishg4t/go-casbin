package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	RedisCache *redis.Client
)

func init() {
	// Connect to DB
	var err error
	DB, err = gorm.Open(postgres.Open("postgres://core:core@localhost:5432/core?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}

	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	RedisCache = redis.NewClient(opt)

	if err != nil {
		panic(fmt.Sprintf("failed to initialize cache: %v", err))
	}
}
