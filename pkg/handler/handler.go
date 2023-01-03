package handler

import (
	_ "app/docs"
	"app/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/registration", h.signUp)
		auth.POST("/login", h.signIn)
		auth.GET("/get-user", h.userIdentity, h.GetUserData)
	}
	transaction := router.Group("/transaction")
	{
		transaction.POST("/add", h.addTransaction)
		transaction.GET("/get-all", h.getAllTransactions)
		transaction.POST("/buy-item", h.userIdentity, h.buyItem)

		nft := transaction.Group("/nft")
		{
			nft.GET("/:id", h.getNftTransactions)
		}
	}
	marketplace := router.Group("/marketplace")
	{
		marketplace.GET("/item-count", h.getItemCount)
		marketplace.GET("/get-all", h.getAllItems)
		marketplace.GET("/my-listings", h.userIdentity, h.getMyListings)
		marketplace.GET("/my-purchases", h.userIdentity, h.getMyPurchases)
		marketplace.GET("/init", h.init)
	}

	return router
}
