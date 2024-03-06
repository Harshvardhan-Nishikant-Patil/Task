// main.go
package main

import (
    "log"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

func main() {
    // Initialize database
    var err error
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate database schema
    err = db.AutoMigrate(&Product{}, &User{}, &CartItem{}, &Order{})
    if err != nil {
        log.Fatal("Failed to auto-migrate database:", err)
    }

    // Initialize Gin router
    router := gin.Default()

    // Enable CORS
    router.Use(cors.Default())

    // Routes
    router.POST("/login", login)

    api := router.Group("/api")
    {
        api.GET("/products", getProducts)
        api.POST("/products", createProduct)
        api.GET("/cart", getCart)
        api.POST("/cart", addToCart)
        api.DELETE("/cart/:id", removeFromCart)
        api.POST("/checkout", checkout)
        // Add more routes for user management, orders, etc.
    }

    // Start server
    err = router.Run(":8080")
    if err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
