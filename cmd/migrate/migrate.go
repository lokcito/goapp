package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/models"
)

func mustEnv(k string) string {
    v := os.Getenv(k)
    if v == "" {
        log.Fatalf("missing env %s", k)
    }
    return v
}

func openDB() *gorm.DB {
    _ = godotenv.Load() // optional .env
    host := mustEnv("DB_HOST")
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "3306"
    }
    user := mustEnv("DB_USER")
    pass := mustEnv("DB_PASS")
    name := mustEnv("DB_NAME")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed connect db: %v", err)
    }
    return db
}

func main() {
    db := openDB()
    // Auto migrate
    if err := db.AutoMigrate(&models.Robot{}); err != nil {
        log.Fatalf("migrate error: %v", err)
    }
    fmt.Println("migrations applied")

    // seed (optional)
    if os.Getenv("SEED") == "1" {
        robots := []models.Robot{
            {Nombre: "Robo-A", Descripcion: "Primer robot"},
            {Nombre: "Robo-B", Descripcion: "Segundo robot"},
            {Nombre: "Robo-C", Descripcion: "Tercer robot"},
        }
        for _, r := range robots {
            db.Create(&r)
        }
        fmt.Println("seed done")
    }
}
