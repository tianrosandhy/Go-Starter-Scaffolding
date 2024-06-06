package bootstrap

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// InitRedis initialize redis connection
func NewRedis(viperConfig *viper.Viper) *redis.Client {
	cli := new(redis.Client)

	redisHost := viperConfig.GetString("REDIS_HOST")
	redisPort := viperConfig.GetString("REDIS_PORT")
	if redisHost != "" {
		cli = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
			Password: viperConfig.GetString("REDIS_PASS"),
			DB:       viper.GetInt("REDIS_INDEX"),
		})

		resp, err := cli.Ping().Result()
		if len(resp) == 0 || err != nil {
			panic("Cannot connect to redis : " + err.Error())
		}
	}

	return cli
}
