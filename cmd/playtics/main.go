package main

import (
	"context"
	"log"

	"playtics/internal/config"
	"playtics/internal/infrastructure/postgres/gen"
	"playtics/internal/registry"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// config
	cfg := config.Load()

	// DB
	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("failed to create connection pool: %v", err)
	}
	defer dbpool.Close()

	// Ping
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// DI
	queries := gen.New(dbpool)
	reg := registry.NewRegistry(queries)

	// Router
	r := gin.Default()
	reg.AppHandler.RegisterRoutes(r)

	// Start
	r.Run()

}
