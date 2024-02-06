package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(incomingRoutes *echo.Echo) {
	incomingRoutes.POST("/login", controllers.Login)
	incomingRoutes.POST("/logout", controllers.Logout)
}
