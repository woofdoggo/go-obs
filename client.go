package go_obs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	connected bool
	ready     bool
	auth      *GetAuthRequiredResponse
	conn      *websocket.Conn
	url       string
	errMap    map[string]chan error
	recvMap   map[string]chan []byte
	mx        sync.Mutex
}

func NewClient(url string) Client {
	return Client{
		false,
		false,
		nil,
		nil,
		url,
		make(map[string]chan error),
		make(map[string]chan []byte),
		sync.Mutex{},
	}
}

func (c *Client) Authenticate(password string) error {
	if !c.connected {
		return errors.New("client not connected")
	}
	if c.auth == nil {
		return errors.New("no auth response")
	}
	saltpwd := password + c.auth.Salt
	salthash := sha256.Sum256([]byte(saltpwd))
	secret := base64.StdEncoding.EncodeToString(salthash[:])
	sec := secret + c.auth.Challenge
	sechash := sha256.Sum256([]byte(sec))
	secRes := base64.StdEncoding.EncodeToString(sechash[:])
	_, err := NewAuthenticateRequest(c, secRes)
	if err != nil {
		return err
	}
	c.ready = true
	return nil
}

func (c *Client) Connect() (bool, chan error, error) {
	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		conn.Close()
		return false, nil, err
	}
	errch := c.poll()
	c.conn = conn
	res, err := NewGetAuthRequiredRequest(c)
	if err != nil {
		conn.Close()
		return false, errch, err
	}
	if res.AuthRequired {
		c.auth = res
	} else {
		c.ready = true
	}
	c.connected = true
	return res.AuthRequired, errch, nil
}

func (c *Client) poll() chan error {
	errch := make(chan error)
	go func() {
		defer func() {
			c.mx.Lock()
			defer c.mx.Unlock()
			c.conn.Close()
			c.connected = false
		}()

		for {
			_, data, err := c.conn.ReadMessage()
			if err != nil {
				errch <- err
				return
			}
			m := make(map[string]interface{})
			err = json.Unmarshal(data, &m)
			if err != nil {
				errch <- err
				return
			}
			if id, ok := m["message-id"]; ok {
				if status, ok := m["status"]; ok {
					if status == "error" {
						errMsg := m["error"]
						c.errMap[id.(string)] <- errors.New(errMsg.(string))
					} else {
						c.recvMap[id.(string)] <- data
					}
				} else {
					errch <- errors.New("no status")
					return
				}
				delete(c.errMap, id.(string))
				delete(c.recvMap, id.(string))
			} else {
				errch <- errors.New("no message-id")
				return
			}
		}
	}()
	return errch
}

func (c *Client) send(data interface{}, id string, errch chan error) chan []byte {
	resch := make(chan []byte)
	if !c.connected {
		errch <- errors.New("client not connected")
		return resch
	}
	if !c.ready {
		errch <- errors.New("client not authenticated")
		return resch
	}
	go func() {
		c.mx.Lock()
		defer c.mx.Unlock()
		c.errMap[id] = errch
		c.recvMap[id] = resch
		err := c.conn.WriteJSON(data)
		if err != nil {
			errch <- err
			delete(c.errMap, id)
			delete(c.recvMap, id)
		}
	}()
	return resch
}
