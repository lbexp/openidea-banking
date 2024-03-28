package config

import "github.com/spf13/viper"

func init() {
	InitEnvBinder()
}

func InitEnvBinder() {
	/**
	 * DB related variables
	 */
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_USERNAME")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_PARAMS")

	/**
	 * AWS S3 related variables
	 */
	viper.BindEnv("S3_ID")
	viper.BindEnv("S3_SECRET_KEY")
	viper.BindEnv("S3_BUCKET_NAME")
	viper.BindEnv("S3_REGION")

	/**
	 * Encryption related variables
	 */
	viper.BindEnv("BCRYPT_SALT")
	viper.BindEnv("JWT_SECRET")
}
