package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/shared/datastore"
	redisx "learnyscape-backend-mono/internal/shared/redis"
	"learnyscape-backend-mono/pkg/database"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var (
	db           *sqlx.DB
	rdb          *redis.Client
	redisClient  redisx.RedisClient
	dataStore    datastore.DataStore
	jwtUtil      jwtutil.JWTUtil
	bcryptHasher encryptutil.Hasher
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	rdb = database.NewRedis((*database.RedisOptions)(cfg.Redis))
	redisClient = redisx.NewRedisClient(rdb)
	dataStore = datastore.NewDataStore(db)
	jwtUtil = jwtutil.NewJWTUtil()
	bcryptHasher = encryptutil.NewBcryptHasher(cfg.App.BCryptCost)
}
