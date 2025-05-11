package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/data"
	"learnyscape-backend-mono/pkg/database"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var (
	db           *sqlx.DB
	rdb          *redis.Client
	dataStore    data.DataStore
	jwtUtil      jwtutil.JWTUtil
	bcryptHasher encryptutil.Hasher
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	rdb = database.NewRedis((*database.RedisOptions)(cfg.Redis))
	dataStore = data.NewDataStore(db)
	jwtUtil = jwtutil.NewJWTUtil()
	bcryptHasher = encryptutil.NewBcryptHasher(cfg.App.BCryptCost)
}
