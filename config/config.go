package config

import (
	"os"
	"regexp"
	"store-dashboard-service/util/log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectName     string
	Port            int
	PostgresDB      DBConfig
	AccessTokenKey  string
	RefreshTokenKey string
	CloudinaryURL   string
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
	ctxScope := "init_config"

	err := godotenv.Load()
	if err != nil {
		// load env for testing
		log.GetLogger().Info(ctxScope, "load env from current working directory", nil)
		re := regexp.MustCompile(`^(.*` + projectName + `)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))
		err = godotenv.Load(string(rootPath) + "/.env.test")
		if err != nil {
			log.GetLogger().Error(ctxScope, "failed to load env", err)
		}
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
		CloudinaryURL:   viper.GetString("CLOUDINARY_URL"),
	}
}

func GetConfig() *Config {
	return &config
}
