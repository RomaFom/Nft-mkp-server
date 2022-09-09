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
	transaction := router.Group("/transaction", h.userIdentity)
	{
		transaction.POST("/add", h.addTransaction)
		transaction.GET("/get-all", h.getAllTransactions)
	}

	return router
}
