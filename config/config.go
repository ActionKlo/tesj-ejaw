package config

import (
	"github.com/ActionKlo/test-ejaw/internal/handlers"
	"github.com/ActionKlo/test-ejaw/internal/repository"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Postgres `mapstructure:",squash"`
}

type Postgres struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DB       string `mapstructure:"POSTGRES_DB"`
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
}

type Services struct {
	H *handlers.RecipesHandler
}

func InitConfig(log *zap.Logger) *Config {
	var appConfig Config
	v := viper.New()
	v.SetConfigType("env")
	v.AddConfigPath(".")    // path for local development
	v.AddConfigPath("/app") // path for container
	v.SetConfigName(".env")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("failed to read config", zap.Error(err))
	}

	if err := v.Unmarshal(&appConfig); err != nil {
		log.Fatal("failed to unmarshal into config struct", zap.Error(err))
	}

	log.Debug("config struct", zap.Any("", appConfig))
	return &appConfig
}

func (cfg *Config) InitServices(logger *zap.Logger) *Services {
	db := repository.OpenDBConnection(logger, repository.DBConfig{
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DB:       cfg.Postgres.DB,
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
	})

	store := repository.InitRepository(logger, db)

	h := handlers.InitRecipesHandler(logger, store)

	return &Services{
		H: h,
	}
}
