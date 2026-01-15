package config

import "github.com/lucas-woo/godotenv"

func InitializeEnv() error {
	return godotenv.LoadEnv()
}