/**
 * @Description: websocket pool
 */
package pool

import (
	"github.com/gorilla/websocket"
)

type websocketItem struct {
	conn *websocket.Conn
	id   string
}

type webSocketPool struct {
	pool []*websocketItem
}

var WebSocketPool *webSocketPool

func init() {
	WebSocketPool = &webSocketPool{
		pool: []*websocketItem{},
	}
}

func (w *webSocketPool) Put(conn *websocket.Conn) {
	item := &websocketItem{
		conn: conn,
		id:   conn.LocalAddr().String(),
	}
	WebSocketPool.pool = append(WebSocketPool.pool, item)
}

func (w *webSocketPool) Seed(task string) {
	for k, v := range w.pool {
		if v.conn != nil {
			err := v.conn.WriteMessage(websocket.TextMessage, []byte(task))
			if err != nil {
				// 如果发生err 就说明连接断开了
				//删除这一个
				v.conn.Close()
				if k != len(w.pool)-1 {
					w.pool = append(w.pool[:k], w.pool[k+1:]...)
				} else {
					w.pool = append(w.pool[:k])
				}
			}
		}

	}
}
