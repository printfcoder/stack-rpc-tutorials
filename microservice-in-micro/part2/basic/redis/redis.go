package redis

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/config"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

// Init 初始化Redis
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("已经初始化过Redis...")
		return
	}

	redisConfig := config.GetRedisConfig()

	// 打开才加载
	if redisConfig != nil && redisConfig.GetEnabled() {
		log.Info("初始化Redis...")

		// 加载哨兵模式
		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			log.Info("初始化Redis，哨兵模式...")
			initSentinel(redisConfig)
		} else { // 普通模式
			log.Info("初始化Redis，普通模式...")
			initSingle(redisConfig)
		}

		log.Info("初始化Redis，检测连接...")

		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Info("初始化Redis，检测连接Ping.")
		log.Info("初始化Redis，检测连接Ping..")
		log.Infof("初始化Redis，检测连接Ping... %s", pong)
	}
	inited = true
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
