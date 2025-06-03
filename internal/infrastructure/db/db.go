package db

import (
    "fmt"

    "cpru-api/ent"
    _ "github.com/lib/pq" // PostgreSQL driver for Ent
)

// NewEntClient สร้างและคืนค่า *ent.Client ที่เชื่อมต่อกับ PostgreSQL
func NewEntClient(dsn string) (*ent.Client, error) {
    // dsn ควรอยู่ในรูปแบบ:
    //   "host=localhost port=5432 user=youruser password=yourpass dbname=yourdb sslmode=disable"
    client, err := ent.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
    }
    return client, nil
}
