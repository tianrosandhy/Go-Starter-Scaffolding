package bootstrap

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/tianrosandhy/goconfigloader"
)

// InitRedis initialize redis connection
func NewRedis(cfg *goconfigloader.Config) *redis.Client {
	cli := new(redis.Client)

	redisHost := cfg.GetString("REDIS_HOST")
	redisPort := cfg.GetString("REDIS_PORT")
	if redisHost != "" {
		cli = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
			Password: cfg.GetString("REDIS_PASS"),
			DB:       viper.GetInt("REDIS_INDEX"),
		})

		resp, err := cli.Ping().Result()
		if len(resp) == 0 || err != nil {
			panic("Cannot connect to redis : " + err.Error())
		}
	}

	return cli
}
