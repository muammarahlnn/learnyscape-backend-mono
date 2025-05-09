package config

import (
	"log"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	AllowedAlgs          []string `mapstructure:"JWT_ALLOWED_ALGS"`
	Issuer               string   `mapstructure:"JWT_ISSUER"`
	AccessSecretKey      string   `mapstructure:"JWT_ACCESS_SECRET_KEY"`
	RefreshSecretKey     string   `mapstructure:"JWT_REFRESH_SECRET_KEY"`
	AccessTokenDuration  int      `mapstructure:"JWT_ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration int      `mapstructure:"JWT_REFRESH_TOKEN_DURATION"`
}

func initJWTConfig() *JWTConfig {
	jwtConfig := &JWTConfig{}

	if err := viper.Unmarshal(&jwtConfig); err != nil {
		log.Fatalf("error mapping jwt config: %v", err)
	}

	return jwtConfig
}
