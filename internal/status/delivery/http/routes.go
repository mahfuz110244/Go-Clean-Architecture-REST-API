package http

import (
	"github.com/labstack/echo/v4"

	"github.com/mahfuz110244/api-mc/internal/middleware"
	"github.com/mahfuz110244/api-mc/internal/status"
)

// Map status routes
func MapStatusRoutes(statusGroup *echo.Group, h status.Handlers, mw *middleware.MiddlewareManager) {
	// statusGroup.POST("/create", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	statusGroup.POST("", h.Create(), mw.AuthSessionMiddleware)
	statusGroup.PUT("/:status_id", h.Update(), mw.AuthSessionMiddleware)
	statusGroup.DELETE("/:status_id", h.Delete(), mw.AuthSessionMiddleware)
	statusGroup.GET("/:status_id", h.GetByID())
	statusGroup.GET("/search", h.SearchByTitle())
	statusGroup.GET("", h.GetStatus())
}
