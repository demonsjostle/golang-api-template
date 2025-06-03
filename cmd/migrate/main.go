package main

import (
	"context"
	"log"
  "fmt"

  "cpru-api/internal/infrastructure/config"
  "cpru-api/internal/infrastructure/db"
  _ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
  cfg := config.LoadConfig()
  dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBSSLMode,
  )
	client, err := db.NewEntClient(dsn) 
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// ไม่ใส่ schema.WithAtlas() แปลว่าใช้ auto migration ธรรมดาของ Ent
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	log.Println("✅ Schema synced with database (auto-migration)")
}
