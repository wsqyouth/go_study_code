package router

import (
	"net/http"
	"todomvc/pkg/todomvc/router/handler/todo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(cors.Default())
	router.StaticFS("/js", http.Dir("./statics/js"))
	router.StaticFS("/node_modules", http.Dir("./statics/node_modules"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	todoRoute := router.Group("/todo")
	todoRoute.GET("", todo.List)
	todoRoute.GET(":id", todo.Get)
	todoRoute.DELETE(":id", todo.Delete)
	todoRoute.POST("", todo.Create)
	todoRoute.PUT(":id", todo.Update)
	todoRoute.POST(":id/done", todo.Toggle)

	return router
}
