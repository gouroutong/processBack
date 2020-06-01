package ws

import (
	"fmt"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)

var ConnIns *WebsocketController

func ConfigureMVC(m *mvc.Application) {
	ws := websocket.New(websocket.Config{})
	// http://localhost:8080/websocket/iris-ws.js
	m.Router.Any("/iris-ws.js", websocket.ClientHandler())

	//这将绑定ws.Upgrade的结果，这是一个websocket.Connection
	//由`m.Handle`服务的控制器。
	m.Register(ws.Upgrade)
	m.Handle(new(WebsocketController))
}

type WebsocketController struct {
	//注意你也可以使用匿名字段，无所谓，binder会找到它。
	//这是当前的websocket连接，每个客户端都有自己的* websocketController实例。
	Conn websocket.Connection
}

func (c *WebsocketController) OnLeave(roomName string) {
	c.Conn.To(websocket.Broadcast).Emit("visit", "leave")
}

func (c *WebsocketController) Update() {
	fmt.Println("visits")
	c.Conn.To(websocket.Broadcast).Emit("visit", "event")
}
func (c *WebsocketController) Get( /* websocket.Connection could be lived here as well, it doesn't matter */) {
	ConnIns = c
	c.Conn.OnLeave(c.OnLeave)
	c.Conn.On("visit", c.Update)
	// 在所有事件回调注册后调用它。
	c.Conn.Wait()
}
