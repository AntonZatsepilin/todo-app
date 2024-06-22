package handler

import (
	"TODO-APP/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() //Создает новый маршрутизатор Gin.

	auth := router.Group("/auth") //Создает группу маршрутов для аутентификации. Группа /auth будет использоваться для всех маршрутов, связанных с аутентификацией.
	{
		auth.POST("/sign-up", h.signUp) //Определяет маршрут POST /sign-up в группе /auth и связывает его с методом signUp, который находится в экземпляре структуры Handler (полученного в качестве параметра функции).
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api") //Создает группу маршрутов для API. Группа /api будет использоваться для всех маршрутов, связанных с API.
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.GetListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.GetItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}

		}
	}
	return router
}
