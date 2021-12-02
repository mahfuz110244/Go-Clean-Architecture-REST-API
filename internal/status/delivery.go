package status

import "github.com/labstack/echo/v4"

// Status HTTP Handlers interface
type Handlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetStatus() echo.HandlerFunc
	SearchByTitle() echo.HandlerFunc
}
