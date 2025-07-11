package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/url"
)

type Client interface {
	Close() error
	Snd(v any) error
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host string

	opt dailOption
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOptions(opts...)

	c := client{
		Conn: nil,
		host: host,
		opt:  opt,
	}
	conn, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = conn
	return &c
}

func (c *client) dail() (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   c.host,
		Path:   c.opt.patten,
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	return conn, err
}

func (c *client) Snd(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = c.WriteMessage(websocket.TextMessage, data)
		if err == nil {
			return nil
		}

		conn, dailErr := c.dail()
		if dailErr != nil {
			return dailErr
		}
		c.Conn = conn
	}

	return err
}

func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(msg, v)
}
