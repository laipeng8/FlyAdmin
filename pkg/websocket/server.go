package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

//调用Example：web.BroadcastMessage(web.Message{Type: websocket.TextMessage, Body: fmt.Sprintf("系统发现有新的订单，目前有%d个订单未处理", total)})

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 在生产环境中应该更严格地检查来源
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// 定义一个广播通道和锁
	broadcast = make(chan Message)
	mu        sync.Mutex
	clients   = make(map[*websocket.Conn]bool)

	// 启动广播goroutine
	wg sync.WaitGroup

	// 心跳设置
	pingPeriod     = (pongWait + writeWait) / 2
	pongWait       = 60 * time.Second
	writeWait      = 10 * time.Second
	maxMessageSize = 512
)

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func init() {
	go broadcastToClients()
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	// 定义最大消息大小为 int 类型
	maxMessageSize = 512

	// 设置读取限制时将其转换为 int64 类型
	conn.SetReadLimit(int64(maxMessageSize))
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// 启动goroutine读取消息
	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadMessages(conn)
	}()

	// 启动心跳检查
	go pingPong(conn)

	// 保持连接，等待客户端断开
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read: %v", err)
			break
		}
	}
}

func ReadMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error: %v", err)
			break
		}
		if err != nil {
			log.Printf("read: %v", err)
			break
		}
		// 处理接收到的消息（这里只是示例）
		log.Printf("received: %s", message)
	}
	RemoveClient(conn)
	conn.Close()
}

func RemoveClient(conn *websocket.Conn) {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
}

func BroadcastMessage(msg Message) {
	broadcast <- msg
}

func broadcastToClients() {
	for {
		select {
		case msg, ok := <-broadcast:
			if !ok {
				log.Println("broadcast channel closed")
				return
			}
			mu.Lock()
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, []byte(msg.Body))
				if err != nil {
					log.Printf("write: %v", err)
					client.Close()
					RemoveClient(client)
				}
			}
			mu.Unlock()
		}
	}
}

func pingPong(conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				log.Println("ping error:", err)
				return
			}
		}
	}
}

func Cleanup() {
	close(broadcast)
	wg.Wait()
}
