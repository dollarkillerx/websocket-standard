/**
 * @Description: websocket server
 */
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"websocketdome/pool"
)

func Websocket(app *gin.Context) {
	value, exists := app.Get("conn")
	if !exists {
		log.Println("not conn")
		return
	}
	conn, ok := value.(*websocket.Conn)
	if !ok {
		log.Println("Transcoding failed")
		return
	}

	// 没有问题将conn放入维护体中
	if conn != nil {
		pool.WebSocketPool.Put(conn)
	}

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(conn.LocalAddr(), "设备断开连接")
			return
		}
		log.Println(string(data))
	}
}
