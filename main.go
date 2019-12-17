/**
 * @Description: websocket demo
 */
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"websocketdome/router"
)

func main() {
	app := gin.Default()

	router.Registered(app)

	log.Println("run 8081")
	if err := app.Run(":8081"); err != nil {
		log.Panic(err)
	}
}
