package main

import (
	"log"
	handler "shopmall/internal/handler/user"
	"shopmall/internal/repository/rpuser"
	"shopmall/internal/service/svuser"
	"shopmall/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	// db, err := initDB(cfg)
	// if err != nil {
	// 	log.Fatalf("Failed to init database: %v", err)
	// }

	// 初始化依赖
	userRepo := rpuser.NewUserMemRepoImpl()
	userService := svuser.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 设置路由
	r := gin.Default()

	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API 路由组
	v1 := r.Group("/api/v1")
	{
		// 用户认证相关接口
		v1.POST("/login", userHandler.Login)

		// 用户管理相关接口
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.GET("/users", userHandler.ListUsers)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// 启动服务器
	if err := r.Run(cfg.Server.HTTP.Addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
