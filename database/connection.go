package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoDB   *mongo.Database
	PostgreDB *pgxpool.Pool
)
