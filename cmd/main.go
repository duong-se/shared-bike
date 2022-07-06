package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	docs "shared-bike/cmd/docs"
	customMiddleware "shared-bike/middleware"
	"shared-bike/pkg/bike"
	"shared-bike/pkg/user"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	swagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

// @title                      Shared Bike API
// @version                    1.0
// @description                This is a shared bike management.
// @contact.name               Duong Pham
// @contact.url                https://github.com/duong-se
// @contact.email              duongpham@duck.com
// @BasePath                   /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	tls := os.Getenv("TLS")
	baseURl := os.Getenv("BASE_URL")
	docs.SwaggerInfo.Schemes = []string{tls}
	docs.SwaggerInfo.Host = baseURl
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if port == "" {
		port = defaultPort
	}
	// Setup
	e := echo.New()
	e.Use(customMiddleware.AddHeaderXRequestID)
	cookieSecret := os.Getenv("COOKIE_SECRET")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(cookieSecret))))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Logger.SetLevel(log.INFO)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	e.GET("/swagger/*", swagger.WrapHandler)
	root := e.Group("/api/v1")
	userRepo := user.NewRepository(db)
	userUseCase := user.NewUseCase(e.Logger, userRepo)
	userHandler := user.NewHandler(userUseCase)
	userAPIs := root.Group("/users")
	userAPIs.POST("/login", userHandler.Login)
	userAPIs.POST("/register", userHandler.Register)

	bikeRepo := bike.NewRepository(db)
	bikeUseCase := bike.NewUseCase(e.Logger, bikeRepo, userRepo)
	bikeHandler := bike.NewHandler(bikeUseCase)
	bikeAPIs := root.Group("/bikes", customMiddleware.Authorize)
	bikeAPIs.GET("", bikeHandler.GetAllBike)
	bikeAPIs.PATCH("/:id/rent", bikeHandler.Rent)
	bikeAPIs.PATCH("/:id/return", bikeHandler.Return)

	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
