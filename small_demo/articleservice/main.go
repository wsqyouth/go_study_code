package main

import (
	"articleservice/entity"

	"articleservice/middleware"
	"articleservice/repository"
	"articleservice/service"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	ar := repository.NewMysqlArticleRepository(dbConn)
	as := service.NewArticleService(ar)

	r := gin.Default() // 已经包含了日志和恢复中间件

	// r.Use(func(c *gin.Context) {
	// 	start := time.Now()
	// 	c.Next()
	// 	log.Printf("-----%s %s %v", c.Request.Method, c.Request.URL, time.Since(start))
	// })
	r.Use(middleware.LoggingMiddleware())
	r.GET("/articles/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		article, err := as.GetArticle(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, article)
	})

	r.POST("/articles", func(c *gin.Context) {
		var article entity.Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := as.CreateArticle(c, &article)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, article)
	})

	log.Fatal(r.Run(":8080"))
}
