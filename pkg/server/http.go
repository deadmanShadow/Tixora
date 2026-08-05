package server

import (
	"fmt"
	"net/http"
	"tixora/pkg/config"
	"tixora/pkg/domain/booking"
	"tixora/pkg/domain/event"
	"tixora/pkg/domain/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func New(db *gorm.DB, cfg *config.Config) *echo.Echo {
	db.AutoMigrate(&user.User{}, &event.Event{}, &booking.Booking{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "running")
	})

	//routes
	user.RegisterRoutes(e, db, cfg)
	event.RegisterRoutes(e, db)
	booking.RegisterRoutes(e, db, cfg)

	return e
}

func Start(db *gorm.DB, cfg *config.Config) {
	e := New(db, cfg)
	port := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}