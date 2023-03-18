package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgreSQLConnection(uri string, connections int32) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}
	config.MaxConns = connections
	//create a new connection pool
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	PostgreDB = pool
	return pool, nil
}
