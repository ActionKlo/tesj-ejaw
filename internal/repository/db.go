package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	db     *sql.DB
}

type DBConfig struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DB       string `mapstructure:"POSTGRES_NAME"`
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
}

func OpenDBConnection(log *zap.Logger, cfg DBConfig) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("failed connect to database", zap.Error(err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal("failed to check ping to database", zap.Error(err))
	}

	log.Info("successfully connected to database")

	return db
}

func InitRepository(logger *zap.Logger, db *sql.DB) *Service {
	return &Service{
		logger: logger,
		db:     db,
	}
}
