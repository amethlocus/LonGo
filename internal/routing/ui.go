package routing

import (
	"LonGo/internal/gintemplrenderer"
	"LonGo/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	router := setupRouter()
	router.LoadHTMLGlob("internal/views/*")

	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	// create connection to DB

	// starts endpoints
	router.GET("/login", handler.Login)
	router.POST("/users", handler.CreateUser)
	router.GET("/users", handler.CreateUserForm)
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
