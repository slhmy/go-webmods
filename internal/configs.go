package internal

const (
	ConfigKeyGORMDatabaseDriver   = "gorm.database.driver"
	ConfigKeyGORMDatabaseName     = "gorm.database.name"
	ConfigKeyGORMDatabaseHost     = "gorm.database.host"
	ConfigKeyGORMDatabasePort     = "gorm.database.port"
	ConfigKeyGORMDatabaseUsername = "gorm.database.username"
	ConfigKeyGORMDatabasePassword = "gorm.database.password"
	ConfigKeyGORMDatabaseSSLMode  = "gorm.database.ssl_mode"

	ConfigKeyRedisAddrs    = "redis.addrs"
	ConfigKeyRedisPassword = "redis.password"

	ConfigKeyMongoHost               = "mongo.host"
	ConfigKeyMongoUsername           = "mongo.username"
	ConfigKeyMongoPassword           = "mongo.password"
	ConfigKeyMongoDatabase           = "mongo.database"
	ConfigKeyMongoSlowQueryThreshold = "mongo.slow_query_threshold"
	ConfigKeyMongoEnableSeedList     = "mongo.enable_seedlist" // https://www.mongodb.com/docs/manual/reference/connection-string/#std-label-connections-dns-seedlist
)
