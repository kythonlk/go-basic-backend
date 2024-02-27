package router

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Home Page",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title": "About Page",
		})
	})

  r.NoRoute(func(c *gin.Context) {
      c.HTML(http.StatusNotFound, "404.html", gin.H{})
  })

	return r
}
