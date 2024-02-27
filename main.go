package main

import (
  "go-cms/router"
)

func main() {

  r := router.Router()

	r.Run(":8000") // Listen and serve on 0.0.0.0:8080
}
