package db

import (
    models "MECHOPS/Models"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB


//connection of database
func Connection() {

    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    root := os.Getenv("DB_ROOT")
    var er error

    DB, er = gorm.Open(mysql.Open(root), &gorm.Config{})
    if er != nil {
        log.Fatal("Failed to connect Database", er)
    }

    err = DB.AutoMigrate(
        &models.User{},
    )
    if err != nil {
        log.Fatal("Failed to AutoMigrate", err)
    }

    log.Println("Database Connected Successfully") 
}
