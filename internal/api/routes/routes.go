package routes

import (
	"lasting-dynamics.com/habit_tracking_app/internal/api/controllers"
	"lasting-dynamics.com/habit_tracking_app/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes for the application.
func initRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api/v1")
	initUserRoutes(apiRouter)
	initHabitRoutes(apiRouter)
}

// initUserRoutes initializes routes related to users.
func initUserRoutes(router *gin.RouterGroup) {
	userHandler := controllers.NewUserHandler()

	// User routes
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/me", middleware.IsAuthenticated(), userHandler.GetUserByID)
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}
}

// initUserRoutes initializes routes related to habits.
func initHabitRoutes(router *gin.RouterGroup) {
	habitHandler := controllers.NewHabitHandler()

	habitRoutes := router.Group("/habits").Use(middleware.IsAuthenticated())
	{
		habitRoutes.GET("/", habitHandler.ListHabits)
		habitRoutes.GET("/:habitId", habitHandler.GetHabitByID)
		habitRoutes.POST("/", habitHandler.CreateHabit)
		habitRoutes.PATCH("/:habitId", habitHandler.UpdateHabit)
		habitRoutes.DELETE("/:habitId", habitHandler.DeleteHabit)
	}
}

// SetupRouter initializes an app router Engine instance with the Logger and Recovery middleware already attached.
// Then inits all app exposed routes for endpoints.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	initRoutes(router)

	return router
}
