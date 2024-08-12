package config

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectName     string
	Port            int
	PostgresDB      DBConfig
	AccessTokenKey  string
	RefreshTokenKey string
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

var config Config

func init() {
	projectName := "store-dashboard-service"

	err := godotenv.Load()
	if err != nil {
		// load env for testing
		re := regexp.MustCompile(`^(.*` + projectName + `)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))
		godotenv.Load(string(rootPath) + "/.env.test")
	}
	viper.AutomaticEnv()

	config = Config{
		ProjectName: projectName,
		Port:        viper.GetInt("PORT"),
		PostgresDB: DBConfig{
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			DBName:   viper.GetString("POSTGRES_DBNAME"),
			SSLMode:  viper.GetString("POSTGRES_SSL_MODE"),
		},
		AccessTokenKey:  viper.GetString("ACCESS_TOKEN_KEY"),
		RefreshTokenKey: viper.GetString("REFRESH_TOKEN_KEY"),
	}
}

func GetConfig() *Config {
	return &config
}
