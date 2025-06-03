package config

import (
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
)

// Config คือ struct ที่เก็บค่าต่างๆ ที่อ่านมาจาก environment
type Config struct {
    // Database
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string

    // JWT
    JWTSecret    string
    JWTExpiresIn time.Duration

    // Server
    Port string

    // (ถ้ามี service อื่นๆ เพิ่มได้)
    // RedisAddr     string
    // RedisPassword string
}

// LoadConfig อ่าน .env แล้วแมปเป็น Config struct
func LoadConfig() *Config {
    // ลองโหลด .env ก่อน ถ้าไม่มีไฟล์ก็ไปอ่านจาก environment จริงๆ
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found or failed to load: %v", err)
    }

    // อ่านค่าแต่ละตัวจาก os.Getenv
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbSSL := os.Getenv("DB_SSLMODE")
    if dbSSL == "" {
        dbSSL = "disable"
    }

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET is required but not set")
    }

    // อ่านเวลาหมดอายุของ JWT (string เช่น "24h", "15m" เป็นต้น)
    jwtExpStr := os.Getenv("JWT_EXPIRES_IN")
    if jwtExpStr == "" {
        jwtExpStr = "24h" // กำหนด default หากไม่ได้ตั้ง
    }
    jwtExpiresIn, err := time.ParseDuration(jwtExpStr)
    if err != nil {
        log.Fatalf("Invalid JWT_EXPIRES_IN value: %v", err)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    return &Config{
        DBHost:       dbHost,
        DBPort:       dbPort,
        DBUser:       dbUser,
        DBPassword:   dbPassword,
        DBName:       dbName,
        DBSSLMode:    dbSSL,
        JWTSecret:    jwtSecret,
        JWTExpiresIn: jwtExpiresIn,
        Port:         port,
        // กรณีมีค่าอื่นๆ ก็อ่านเพิ่มได้ที่นี่
        // RedisAddr:     os.Getenv("REDIS_ADDR"),
        // RedisPassword: os.Getenv("REDIS_PASSWORD"),
    }
}
