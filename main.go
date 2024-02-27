package main

import "github.com/gin-gonic/gin"
import "fmt"
import "net/http"

func main() {
  fmt.Println("app running in http://0.0.0.0:8080")
	r := gin.Default()
  
  r.LoadHTMLGlob("src/ui/*.templ")

r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.templ", gin.H{
      "title": "Home",
    })
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.templ", gin.H{})
	})

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
