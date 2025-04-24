// Package env list all environment variables key in the app
package env

const (
	// RedisAddr is the address (including protocol and port) of the Redis server
	RedisAddr = "REDIS_ADDR"
	// RedisPassword is the password for authenticating with Redis
	RedisPassword = "REDIS_PASSWORD"
	// RedisUsername is the username for Redis authentication
	RedisUsername = "REDIS_USERNAME"
	// RedisTLSEnabled indicates whether TLS should be enabled for the Redis connection
	RedisTLSEnabled = "REDIS_TLS_ENABLED"
	// RedisDB is the db (int)
	RedisDB = "REDIS_DB"
)
