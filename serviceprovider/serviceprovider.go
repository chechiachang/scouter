package serviceprovider

import (
	"github.com/linkernetworks/logger"

	"github.com/linkernetworks/mongo"
	"github.com/linkernetworks/redis"
)

type Container struct {
	Config Config
	Redis  *redis.Service
	Mongo  *mongo.Service
}

// Service is the interface
type Service interface{}

// New will create container
func New(cf Config) *Container {
	// setup logger configuration
	logger.Setup(cf.Logger)

	logger.Infof("Connecting to redis: %s", cf.Redis.Addr())
	redisService := redis.New(cf.Redis)

	logger.Infof("Connecting to mongodb: %s", cf.Mongo.Url)
	mongo := mongo.New(cf.Mongo.Url)

	sp := &Container{
		Config: cf,
		Redis:  redisService,
		Mongo:  mongo,
	}

	return sp
}

// NewForTesting will test container for creating a container
func NewForTesting(cf Config) *Container {
	// setup logger configuration
	logger.Setup(cf.Logger)

	logger.Infof("Connecting to redis: %s", cf.Redis.Addr())
	redisService := redis.New(cf.Redis)

	logger.Infof("Connecting to mongodb: %s", cf.Mongo.Url)
	mongo := mongo.New(cf.Mongo.Url)

	sp := &Container{
		Config: cf,
		Redis:  redisService,
		Mongo:  mongo,
	}

	return sp
}
