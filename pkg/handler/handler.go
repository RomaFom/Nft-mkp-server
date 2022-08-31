package handler

import (
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/get-user-data", h.userIdentity, h.GetUserData)
	}
	//api := router.Group("/api", h.userIdentity)
	//{
	//	lists := api.Group("/lists")
	//	{
	//		//GET
	//		lists.GET("/", h.getAllLists)
	//		lists.GET("/:id", h.getListById)
	//		lists.POST("/", h.createList)
	//		lists.PUT("/:id", h.updateList)
	//		lists.DELETE("/:id", h.deleteList)
	//		items := lists.Group(":id/items")
	//		{
	//			items.GET("/", h.getAllItems)
	//			items.POST("/", h.createItem)
	//		}
	//	}
	//	items := api.Group("/item")
	//	{
	//		items.GET("/:id", h.getItemById)
	//		items.PUT("/:id", h.updateItem)
	//		items.DELETE("/:id", h.deleteItem)
	//	}
	//}
	return router
}
