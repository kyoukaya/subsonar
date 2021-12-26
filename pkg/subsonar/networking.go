package subsonar

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack"

	"github.com/kyoukaya/subsonar/internal/util"
)

func (c *Client) send(v Opcoder) error {
	b, err := msgpack.Marshal(&MsgWrapper{
		Opcode: v.Opcode(),
		V:      v,
	})
	if err != nil {
		return err
	}
	// Unlike server messages, client to sever messages aren't compressed
	log.Println(">", opToString[v.Opcode()], util.JSONDump(v))
	return c.w.WriteMessage(websocket.BinaryMessage, b)
}

var ErrTextMessage = fmt.Errorf("unexpected text message")

func (c *Client) read() ([]byte, error) {
	mt, compressed, err := c.w.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("read message: %w", err)
	}
	if mt != websocket.BinaryMessage {
		return nil, fmt.Errorf("message type %d: %w", mt, ErrTextMessage)
	}
	b, err := io.ReadAll(brotli.NewReader(bytes.NewBuffer(compressed)))
	if err != nil {
		return nil, fmt.Errorf("decompress: %w", err)
	}
	return b, nil
}
