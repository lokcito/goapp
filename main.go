package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/controllers"
    "equisd.com/bichito/robotapp/models"
)

func mustGetEnv(key string) string {
    v := os.Getenv(key)
    if v == "" {
        log.Fatalf("missing env %s", key)
    }
    return v
}

func openDB() *gorm.DB {
    host := mustGetEnv("DB_HOST")
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "3306"
    }
    user := mustGetEnv("DB_USER")
    pass := mustGetEnv("DB_PASS")
    name := mustGetEnv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed connect db: %v", err)
    }
    return db
}

func main() {
    // load minimal env defaults
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }
    gin.SetMode(gin.DebugMode)
    
    db := openDB()

    // pass DB to models/controllers
    models.SetDB(db)

    r := gin.Default()

    // load html templates (use layout)
    r.SetFuncMap(template.FuncMap{
        "nl2br": func(s string) template.HTML {
            return template.HTML(template.HTMLEscapeString(s))
        },
    })
    r.LoadHTMLGlob("templates/*.html")
    r.Static("/static", "./static")

    // routes
    r.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/robots")
    })

    robotController := controllers.NewRobotController(db)
    r.GET("/robots", robotController.Index)
    r.GET("/robots/new", robotController.New)
    r.POST("/robots", robotController.Create)
    r.GET("/robots/show/:id", robotController.Show)
    r.GET("/robots/edit/:id", robotController.Edit)
    r.POST("/robots/update/:id", robotController.Update)
    r.POST("/robots/delete/:id", robotController.Delete)

    log.Printf("listening on :%s\n", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
