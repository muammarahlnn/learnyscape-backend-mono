package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	JwtConfig  *JWTConfig
	SmtpConfig *SMTPConfig
)

func init() {
	configPath := parseConfigPath()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	JwtConfig = initJWTConfig()
	SmtpConfig = initSMTPConfig()
}

func parseConfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd)
}
