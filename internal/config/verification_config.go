package config

import (
	"log"

	"github.com/spf13/viper"
)

type VerificationConfig struct {
	AccountVerificationTokenDuration         int `mapstructure:"ACCOUNT_VERIFICATION_TOKEN_DURATION"`
	AccountVerificationTokenCooldownDuration int `mapstructure:"ACCOUNT_VERIFICATION_TOKEN_COOLDOWN_DURATION"`
}

func initVerificationConfig() *VerificationConfig {
	verificationConfig := &VerificationConfig{}
	if err := viper.Unmarshal(&verificationConfig); err != nil {
		log.Fatalf("error mapping verification config: %v", err)
	}

	return verificationConfig
}
