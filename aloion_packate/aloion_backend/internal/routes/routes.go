package routes

import (
	"github.com/arnob17/aloion_backend/internal/handlers"
	"github.com/arnob17/aloion_backend/internal/middleware"
	"github.com/arnob17/aloion_backend/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, jwtSecret string, jwtExpiresIn int) {
	// Initialize services
	userService := services.NewUserService()
	courseService := services.NewCourseService()
	examService := services.NewExamService()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService, jwtSecret, jwtExpiresIn)
	courseHandler := handlers.NewCourseHandler(courseService)
	examHandler := handlers.NewExamHandler(examService)

	// Public routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "message": "Aloion Backend is running"})
		})

		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}

		// Public course routes
		courses := api.Group("/courses")
		{
			courses.GET("", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)
		}

		// Public exam routes
		exams := api.Group("/exams")
		{
			exams.GET("", examHandler.GetAllExams)
			exams.GET("/:id", examHandler.GetExam)
		}
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		users := protected.Group("/users")
		{
			users.GET("/me", userHandler.GetProfile)
			users.GET("", middleware.AdminOnly(), userHandler.GetAllUsers)
		}

		// Course management routes
		courses := protected.Group("/courses")
		{
			courses.POST("", middleware.TeacherOrAdmin(), courseHandler.Create)
			courses.PUT("/:id", middleware.TeacherOrAdmin(), courseHandler.Update)
			courses.DELETE("/:id", middleware.TeacherOrAdmin(), courseHandler.Delete)
			courses.POST("/:id/enroll", courseHandler.Enroll)
			courses.GET("/my-enrollments", courseHandler.GetMyEnrollments)
		}

		// Exam routes
		exams := protected.Group("/exams")
		{
			exams.POST("", middleware.TeacherOrAdmin(), examHandler.CreateExam)
			exams.GET("/my-results", examHandler.GetMyResults)
			exams.GET("/results/:id", examHandler.GetResult)
			exams.POST("/results", middleware.TeacherOrAdmin(), examHandler.AddResult)
			exams.GET("/results", middleware.TeacherOrAdmin(), examHandler.GetAllResults)
		}
	}
}
