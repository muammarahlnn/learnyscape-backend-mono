package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/shared/datastore"
	redisx "learnyscape-backend-mono/internal/shared/redis"
	"learnyscape-backend-mono/pkg/database"
	"learnyscape-backend-mono/pkg/mq"
	encryptutil "learnyscape-backend-mono/pkg/util/encrypt"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"
	smtputil "learnyscape-backend-mono/pkg/util/smtp"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var (
	db           *sqlx.DB
	rdb          *redis.Client
	rabbitmq     *amqp.Connection
	redisClient  redisx.RedisClient
	dataStore    datastore.DataStore
	jwtUtil      jwtutil.JWTUtil
	bcryptHasher encryptutil.Hasher
	mailer       smtputil.Mailer
)

func BootstrapGlobal(cfg *config.Config) {
	db = database.NewPostgres((*database.PostgresOptions)(cfg.Postgres))
	rdb = database.NewRedis((*database.RedisOptions)(cfg.Redis))
	rabbitmq = mq.NewAMQP((*mq.AMQPOptions)(cfg.Amqp))
	redisClient = redisx.NewRedisClient(rdb)
	dataStore = datastore.NewDataStore(db)
	jwtUtil = jwtutil.NewJWTUtil()
	bcryptHasher = encryptutil.NewBcryptHasher(cfg.App.BCryptCost)
	mailer = smtputil.NewMailer()
}
