package router
import (
	"hostel-saas/config"
	"hostel-saas/handlers"
	"hostel-saas/middlewares"
	"hostel-saas/pkg/db"
	"hostel-saas/repositories"
	"hostel-saas/services"
	"github.com/gin-gonic/gin"
)
func SetupRouter(cfg *config.Config) *gin.Engine {
	if cfg.Env == "production" { gin.SetMode(gin.ReleaseMode) }
	r := gin.New()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RateLimiter())
	r.Use(middlewares.Logger())
	r.Use(middlewares.Recovery())

	authRepo := repositories.NewAuthRepository(db.DB)
	authService := services.NewAuthService(authRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	orgRepo := repositories.NewOrgRepository(db.DB)
	orgService := services.NewOrgService(orgRepo)
	orgHandler := handlers.NewOrgHandler(orgService)

	hostelRepo := repositories.NewHostelRepository(db.DB)
	hostelService := services.NewHostelService(hostelRepo)
	hostelHandler := handlers.NewHostelHandler(hostelService)

	studentRepo := repositories.NewStudentRepository(db.DB)
	studentService := services.NewStudentService(studentRepo, hostelRepo)
	studentHandler := handlers.NewStudentHandler(studentService)

	financeRepo := repositories.NewFinanceRepository(db.DB)
	financeService := services.NewFinanceService(financeRepo)
	financeHandler := handlers.NewFinanceHandler(financeService)

	analyticsRepo := repositories.NewAnalyticsRepository(db.DB)
	analyticsService := services.NewAnalyticsService(analyticsRepo)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/signup", authHandler.Signup)
		}

		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware(cfg))
		{
			protected.GET("/me", func(c *gin.Context) {
				userID, _ := c.Get("userID")
				orgID, _ := c.Get("orgID")
				role, _ := c.Get("role")
				c.JSON(200, gin.H{"user_id": userID, "org_id": orgID, "role": role})
			})

			org := protected.Group("/organization")
			{
				org.GET("/", orgHandler.GetOrganization)
				org.PUT("/", middlewares.RoleMiddleware("Owner", "Admin"), orgHandler.UpdateOrganization)

				users := org.Group("/users")
				users.Use(middlewares.RoleMiddleware("Owner", "Admin"))
				{
					users.POST("/", orgHandler.CreateUser)
					users.GET("/", orgHandler.ListUsers)
					users.PUT("/:id", orgHandler.UpdateUser)
				}
			}

			hostels := protected.Group("/hostels")
			{
				hostels.GET("/", hostelHandler.ListHostels)
				hostels.POST("/", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), hostelHandler.CreateHostel)
				hostels.GET("/:id", hostelHandler.GetHostelHierarchy)

				hostels.POST("/:id/floors", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), hostelHandler.CreateFloor)
				hostels.POST("/floors/:floor_id/rooms", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), hostelHandler.CreateRoom)
				hostels.POST("/rooms/:room_id/beds", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), hostelHandler.CreateBed)
			}

			students := protected.Group("/students")
			{
				students.POST("/", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), studentHandler.CreateStudent)
				students.GET("/", studentHandler.ListStudents)
				students.GET("/:id", studentHandler.GetStudent)
				students.PUT("/:id", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), studentHandler.UpdateStudent)

				students.POST("/:id/check-in", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), studentHandler.CheckIn)
				students.POST("/:id/check-out", middlewares.RoleMiddleware("Owner", "Admin", "Manager"), studentHandler.CheckOut)
			}

			finance := protected.Group("/finance")
			{
				finance.POST("/rent", middlewares.RoleMiddleware("Owner", "Admin", "Manager", "Accountant"), financeHandler.GenerateRent)
				finance.GET("/rent", financeHandler.ListRents)
				finance.POST("/rent/:rent_id/payment", middlewares.RoleMiddleware("Owner", "Admin", "Manager", "Accountant"), financeHandler.ProcessPayment)
			}

			analytics := protected.Group("/analytics")
			{
				analytics.GET("/kpis", analyticsHandler.GetDashboardKPIs)
			}
		}
	}
	return r
}
