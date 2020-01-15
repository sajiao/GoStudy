package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512

	authToken = "123456"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	ws       *websocket.Conn
	send     chan []byte
	auth     bool
	username []byte
}

func (c *connection) readPump() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		mtype := 2
		text := string(message)
		reg := regexp.MustCompile(`=[^&]+`)
		s := reg.FindAllString(text, -1)

		if len(s) == 2 {
			fromuser := strings.Replace(s[0], "=", "", 1)
			token := strings.Replace(s[1], "=", "", 1)
			if token == authToken {
				c.username = []byte(fromuser)
				c.auth = true
				message = []byte(fromuser + "join")
				mtype = 1
			}
		}

		touser := []byte("all")
		reg2 := regexp.MustCompile(`^@.*? `)
		s2 := reg2.FindAllString(text, -1)
		if len(s2) == 1 {
			s2[0] = strings.Replace(s2[0], "@", "", 1)
			s2[0] = strings.Replace(s2[0], " ", "", 1)
			touser = []byte(s2[0])
		}
		if c.auth == true {
			t := time.Now().Unix()
			h.broadcast <- &tmessage{content: message, fromuser: c.username, touser: touser, mtype: mtype, createtime: time.Unix(t, 0).String()}
		}
	}
}

//给消息，指定消息类型和荷载
// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

func (c *connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &connection{send: make(chan []byte, 256), ws: ws, auth: false}
	h.register <- c
	go c.writePump()
	c.readPump()
}
