/**
 * @Description: task release
 */
package server

import (
	"github.com/gin-gonic/gin"
	"websocketdome/pool"
)

func TaskRelease(app *gin.Context) {
	task := app.PostForm("task")
	pool.WebSocketPool.Seed(task)
	app.JSON(200,"200OK")
}
