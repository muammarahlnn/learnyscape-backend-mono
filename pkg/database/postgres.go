package database

import (
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresOptions struct {
	Host            string
	DBName          string
	Username        string
	Password        string
	SSLMode         string
	Port            int
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime int
}

func NewPostgres(opt *PostgresOptions) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		opt.Host,
		opt.Username,
		opt.Password,
		opt.DBName,
		opt.Port,
		opt.SSLMode,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging to database %v", err)
	}

	db.SetMaxIdleConns(opt.MaxIdleConn)
	db.SetMaxOpenConns(opt.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(opt.MaxConnLifetime) * time.Minute)

	return db
}
