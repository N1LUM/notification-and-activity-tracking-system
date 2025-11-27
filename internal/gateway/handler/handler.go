package handler

import (
	"notify-activity-tracking-system/internal/gateway/auth"
	"notify-activity-tracking-system/internal/gateway/middleware"
	"notify-activity-tracking-system/internal/gateway/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *service.Service
	jwtManager *auth.JWTManager
}

func NewHandler(services *service.Service, jwtManager *auth.JWTManager) *Handler {
	return &Handler{
		services:   services,
		jwtManager: jwtManager,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	authorization := router.Group("/api/auth")
	{
		authorization.POST("/login", h.Login)     // Логин
		authorization.POST("/refresh", h.Refresh) // Рефреш токена
	}

	usersOpen := router.Group("/api/users")
	{
		usersOpen.POST("/", h.CreateUser) // Регистрация
	}

	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware(h.jwtManager))
	{
		users := api.Group("/users")
		{
			users.GET("/", h.GetUsers)                               // Получить список пользователей
			users.GET("/:id", h.GetUserByID)                         // Получить пользователя по ID
			users.PATCH("/:id", h.UpdateUser)                        // Обновить пользователя по ID
			users.DELETE("/:id", h.DeleteUser)                       // Удалить пользователя по ID
			users.GET("/by-username/:username", h.GetUserByUsername) // Получить пользователя по username
		}
		events := api.Group("/events")
		{
			events.POST("/", h.PostEvents) // Отправить ивенты
		}
	}

	return router
}
