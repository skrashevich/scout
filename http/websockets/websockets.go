package websockets

import (
	"sync"

	websocket "github.com/gofiber/websocket/v2"
)

// Run is a wrapper for gofiber websocket
//   socketClosed - passes out to callers that the socket has closed
//   receive - do not block as it will continuously run until the websocket is closed
//   send - push out msgs and return when server is done with the websocket
//   cleanup - will be called when all goroutines are finished
func Run(c *websocket.Conn, socketClosed chan bool, receive func(int, []byte), send func(*websocket.Conn), cleanup func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Read goroutine will cleanup after websocket closes no need to wait for it
	go func() {
		defer close(socketClosed)
		for {
			msgType, data, err := c.ReadMessage()
			if err != nil {
				// socket closed
				return
			}
			if receive != nil {
				receive(msgType, data)
			}
		}
	}()
	go func() {
		defer wg.Done()
		if send != nil {
			send(c)
		}
	}()
	wg.Wait()
	if cleanup != nil {
		cleanup()
	}
}
