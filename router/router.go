/**
 * @Description: 路由层
 */
package router

import (
	"github.com/gin-gonic/gin"
	"websocketdome/middleware"
	"websocketdome/server"
)

func Registered(app *gin.Engine) {
	app.GET("/websocket", middleware.BaseMiddleware, server.Websocket) // 维护websocket相关的连接
	app.POST("/task_release", server.TaskRelease)                      // 下发任务
}
