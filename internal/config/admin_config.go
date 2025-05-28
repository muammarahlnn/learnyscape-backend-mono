package config

import (
	"log"

	"github.com/spf13/viper"
)

type AdminConfig struct {
	AccountVerificationTokenDuration int `mapstructure:"ACCOUNT_VERIFICATION_TOKEN_DURATION"`
}

func initAdminConfig() *AdminConfig {
	adminConfig := &AdminConfig{}
	if err := viper.Unmarshal(&adminConfig); err != nil {
		log.Fatalf("error mapping admin config: %v", err)
	}

	return adminConfig
}
