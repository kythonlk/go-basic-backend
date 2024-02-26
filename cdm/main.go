// package main
//
// import "github.com/gin-gonic/gin"
//
// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"hello": "world",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }
package main

import (
	"context"
	"os"
)

func main() {
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}ackage main

import (
	"context"
	"os"
)

func main() {
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}
