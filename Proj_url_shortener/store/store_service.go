package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//Define the struct wrapper around raw Redis Client

type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeServices and redis context

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

//Note that  in real world usage,the cache duration shouldn't have
//an expiration tim ,an lru Policy config shoulld br set where the
//values that are retrevied less often are purged automatically from
//the cache and stored back in RDBMS whenever the cache is full

const CacheDuration = 6 * time.Hour

// Intializing the store service and return a store pointer

func Intializestore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:8182",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v ,err"))
	}
	fmt.Printf("/nRedis started sucessfully :pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/*
   we want to be able to save theh mapping between  the orginalUrl an the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string, orginalUrl string, userrId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, orginalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("failed to saving key Url | Error:%v -shortUrl:%s - originalUrl:%s\n", err, shortUrl, orginalUrl))

	}
}

/*
we should be able to retrive the intial long URL once the short is provided.
This is when users will calling the shortlink in the url ,so what we need to do
here is to retrive the long url and think about redirect
*/

func RetrieveIntialUrl(shortUrl string) string {
	result, err := storeService.redisClient.GET(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("failed RetrieveIntial Url | Error:%v -shortUrl:%s\n", err, shortUrl))
	}
	return result

}
