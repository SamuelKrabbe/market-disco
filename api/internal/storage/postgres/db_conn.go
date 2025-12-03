package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/SamuelKrabbe/market-disco/api/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPsqlDB(ctx context.Context, c *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)

	// Config do pool
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = maxOpenConns
	cfg.MinConns = maxIdleConns
	cfg.MaxConnIdleTime = connMaxIdleTime * time.Second
	cfg.MaxConnLifetime = connMaxLifetime * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	fmt.Println("connected to database")

	return pool, nil
}
