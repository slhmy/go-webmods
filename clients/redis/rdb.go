package gwm_redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
	gwm_app "github.com/slhmy/go-webmods/app"
	"github.com/slhmy/go-webmods/internal"
)

var (
	rdbInitMutex sync.Mutex
	rdbClient    redis.UniversalClient
)

func GetRDB() redis.UniversalClient {
	if rdbClient == nil {
		rdbInitMutex.Lock()
		defer rdbInitMutex.Unlock()
		if rdbClient != nil {
			return rdbClient
		}

		addrs := gwm_app.Config().GetStringSlice(internal.ConfigKeyRedisAddrs)
		password := gwm_app.Config().GetString(internal.ConfigKeyRedisPassword)
		if len(addrs) == 0 {
			panic("No redis hosts configured")
		}
		if len(addrs) == 1 {
			rdbClient = redis.NewClient(&redis.Options{
				Addr:     addrs[0],
				Password: password,
			})
		}
		if len(addrs) > 1 {
			rdbClient = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:    addrs,
				Password: password,
			})
		}
	}
	return rdbClient
}
