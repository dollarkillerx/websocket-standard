/**
 * @Description: 验证层
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func BaseMiddleware(app *gin.Context) {
	conn, e := upgrader.Upgrade(app.Writer, app.Request, nil)
	if e != nil {
		log.Println("sc: ", e)
		app.Abort()
		return
	}
	app.Set("conn", conn)
	app.Next()
}
