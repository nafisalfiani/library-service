package main

import (
	"fmt"
	"library/config"
	"library/docs"
	"library/entity"
	"library/handler"
	"library/repository"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/xendit/xendit-go/v4"
)

// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// init config
	cfg, err := config.InitEnv()
	if err != nil {
		log.Fatalln(err)
	}

	// init logger
	logger, err := config.InitLogger(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// init DB connection
	db, err := config.InitSql(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// auto migrate DB changes
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Book{},
		&entity.Rental{},
		&entity.RentalDetail{},
		&entity.Payment{},
	); err != nil {
		log.Fatalln(err)
	}

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init xendit
	xnd := xendit.NewClient(cfg.Xendit.ApiKey)

	// init repository
	repo := repository.InitRepository(db, xnd)

	// init handler
	handler := handler.Init(cfg, repo, validator, logger)

	// init echo instance
	e := echo.New()

	// e.Use(middleware.Recover())
	e.Use(handler.MiddlewareLogging)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	docs.SwaggerInfo.Title = "Library Service API"
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.GET("/ping", handler.Ping)

	users := e.Group("/users")
	users.POST("/register", handler.Register)
	users.POST("/login", handler.Login)
	users.GET("", handler.GetUser, handler.Authorize)

	deposits := e.Group("/deposits", handler.Authorize)
	deposits.POST("", handler.TopUpDeposit)

	rentals := e.Group("/rentals", handler.Authorize)
	rentals.GET("/active", handler.GetActiveRental)
	rentals.GET("/closed", handler.GetRentalHistory)
	rentals.POST("", handler.CreateRental)

	payments := e.Group("/payments", handler.Authorize)
	payments.GET("", handler.GetPayments)
	payments.POST("", handler.RefreshPaymentStatus)

	books := e.Group("/books", handler.Authorize)
	books.GET("", handler.ListBook)
	books.GET("/:id", handler.GetBook)
	books.POST("", handler.CreateBook)
	books.PUT("/:id", handler.UpdateBook)
	books.DELETE("/:id", handler.DeleteBook)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", cfg.Server.Base, cfg.Server.Port)))
}
