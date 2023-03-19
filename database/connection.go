package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	PostgreDB *pgxpool.Pool
)
