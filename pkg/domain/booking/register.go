package booking

import (
	"tixora/pkg/auth"
	"tixora/pkg/config"
	"tixora/pkg/domain/event"
	"tixora/pkg/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	bookingRepo := NewRepository(db)
	eventRepo := event.NewRepository(db)

	svc := NewService(bookingRepo, eventRepo)
	handler := NewHandler(svc)

	jwtService := auth.NewJWTService(cfg.JwtSecret)

	api := e.Group("/api/v1/bookings", middlewares.AuthMiddleware(jwtService))

	api.POST("", handler.CreateBooking)
	api.GET("/me", handler.GetMyBookings)
	api.PATCH("/:id/cancel", handler.CancelBooking)

}