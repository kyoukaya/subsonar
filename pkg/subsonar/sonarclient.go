package subsonar

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/kyoukaya/subsonar/internal/util"
)

type Client struct {
	w Socket
}

type Socket interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
}

func New(w Socket) *Client {
	return &Client{w: w}
}

var ErrNotImplemented = errors.New("not implemented")

func (c *Client) Run() {
	for {
		b, err := c.read()
		if err == ErrTextMessage {
			log.Println("skipping unexpected text message", err)
			continue
		}
		if err != nil {
			log.Println("!up readerr:", err)
			return
		}
		op, v, err := unpack(b)
		if err != nil {
			s := base64.RawStdEncoding.EncodeToString(b)
			log.Printf("!up read %d: %s. failed to unpack msg: %v", op, s, err)
			return
		}
		log.Println("<", opToString[op])
		if err := c.Handle(v); err != nil {
			log.Printf("!op code %d %s is unhandled", op, util.JSONDump(v))
		}
	}
}

func (c *Client) Handle(v interface{}) error {
	switch m := v.(type) {
	case *TimeSyncMessage:
		return c.TimeSyncMessageHandler(m)
	case *PingMessage:
		return c.PingMessageHandler(m)
	case *LogMessage:
		return c.LogMessageHandler(m)
	case *SonarMessage:
		return c.SonarMessageHandler(m)
	case *ServerReadyMessage:
		return c.SeverReadyMessageHandler(m)
	case *MessageList:
		for _, v1 := range *m {
			err := c.Handle(v1)
			if err != nil {
				return err
			}
		}
	case *HuntRelayStateMessage:
		// TODO: track and expose the underlying state of hunts
	case *FateRelayStateMessage:
		// TODO: track and expose the underlying state of hunts
	default:
		return ErrNotImplemented
	}
	return nil
}
