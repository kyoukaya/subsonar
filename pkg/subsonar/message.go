package subsonar

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/vmihailenco/msgpack"
)

type Opcoder interface {
	Opcode() uint8
}

type Vec3 struct {
	X int32
	Y int32
	Z int32
}

type ReleaseMode uint8

const (
	ReleaseModeNormal ReleaseMode = iota
	ReleaseModeHold
	ReleaseModeForced
)

var opToType = map[uint8]interface{}{
	0: ServerReadyMessage{},
	// 1: LogMessage{},
	// 2:   SonarMessage{},
	// 16:  PlayerInfoMessage{},
	// 17: PlayerPlaceMessage{},
	// 18: SonarConfig{},
	// 19: SonarVersionMessage{},
	// 20:  "ClientModifiers",
	// 32:  "AuthLogin",
	// 33:  "AuthRegister",
	// 34:  "AuthChange",
	// 35:  "VerificationRequest",
	// 48:  "HuntRelay",
	// 49:  "FateRelay",
	// 63:  "ManualRelay",
	64: HuntRelayStateMessage{},
	65: FateRelayStateMessage{},
	// 128: "RelayDataRequest",
	208: PingMessage{},
	// 209: "PongMessage",
	// 224: "TimeSyncMessageOld",
	225: TimeSyncMessage{},
	255: MessageList{},
}

var opToString = map[uint8]string{
	0:   "ServerReady",
	1:   "LogMessage",
	2:   "SonarMessage",
	16:  "PlayerInfo",
	17:  "PlayerPlace",
	18:  "SonarConfig",
	19:  "SonarVersion",
	20:  "ClientModifiers",
	32:  "AuthLogin",
	33:  "AuthRegister",
	34:  "AuthChange",
	35:  "VerificationRequest",
	48:  "HuntRelay",
	49:  "FateRelay",
	63:  "ManualRelay",
	64:  "RelayState<HuntRelay>",
	65:  "RelayState<FateRelay>",
	128: "RelayDataRequest",
	208: "PingMessage",
	209: "PongMessage",
	224: "TimeSyncMessageOld",
	225: "TimeSyncMessage",
	255: "MessageList",
}

type MsgWrapper struct {
	Opcode uint8
	V      interface{}
}

func (m *MsgWrapper) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var l int
	if l, err = d.DecodeArrayLen(); err != nil {
		return err
	}
	if l != 2 {
		return fmt.Errorf("array len doesn't match: %d", l)
	}
	if m.Opcode, err = d.DecodeUint8(); err != nil {
		return err
	}
	v, ok := opToType[m.Opcode]
	if !ok {
		m.V, err = d.DecodeInterface()
		return fmt.Errorf("failed to find opcode: %w", err)
	}
	i := reflect.New(reflect.TypeOf(v)).Interface()
	if err := d.Decode(i); err != nil {
		return err
	}
	m.V = i
	return nil
}

func (m *MsgWrapper) EncodeMsgpack(e *msgpack.Encoder) error {
	err := e.EncodeArrayLen(2)
	if err != nil {
		return err
	}
	err = e.EncodeUint8(m.Opcode)
	if err != nil {
		return err
	}
	return e.Encode(m.V)
}

func unpack(b []byte) (opcode uint8, v interface{}, err error) {
	m := &MsgWrapper{}
	d := msgpack.NewDecoder(bytes.NewBuffer(b))
	err = m.DecodeMsgpack(d)
	if err != nil {
		return m.Opcode, m.V, err
	}
	return m.Opcode, m.V, nil
}
