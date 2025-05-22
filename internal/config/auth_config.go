package config

import (
	"log"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	RefreshTokenDuration             int `mapstructure:"JWT_REFRESH_TOKEN_DURATION"`
	AccountVerificationTokenDuration int `mapstructure:"ACCOUNT_VERIFICATION_TOKEN_DURATION"`
}

func initAuthConfig() *AuthConfig {
	authConfig := &AuthConfig{}
	if err := viper.Unmarshal(&authConfig); err != nil {
		log.Fatalf("error mapping auth config: %v", err)
	}

	return authConfig
}
