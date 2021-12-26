package subsonar

import (
	"fmt"
	"time"

	"github.com/vmihailenco/msgpack"
)

type TimeSyncMessage struct {
	UnixTime time.Time
}

func (m *TimeSyncMessage) DecodeMsgpack(d *msgpack.Decoder) error {
	n, err := d.DecodeArrayLen()
	if err != nil {
		return err
	}
	if n != 1 {
		return fmt.Errorf("unexpected len of TimeSyncMessage: %d", n)
	}
	v, err := d.DecodeFloat64()
	if err != nil {
		return err
	}
	m.UnixTime = time.UnixMilli(int64(v))
	return nil
}

type ServerReadyMessage struct {
	ConnectionID uint32
}

func (m *ServerReadyMessage) DecodeMsgpack(d *msgpack.Decoder) error {
	n, err := d.DecodeArrayLen()
	if err != nil {
		return err
	}
	if n != 1 {
		return fmt.Errorf("unexpected len of ServerReadyMessage: %d", n)
	}
	v, err := d.DecodeUint32()
	if err != nil {
		return err
	}
	m.ConnectionID = v
	return nil
}

type LogMessage struct{}   // TODO:
type SonarMessage struct{} // TODO:

type PingMessage struct{} // TODO:

type HuntRelayStateMessage struct {
	A struct {
		WorldID     uint32
		ZoneID      uint32
		InstanceID  uint32
		Coords      Vec3
		ID          uint32
		ReleaseMode ReleaseMode
		ActorID     uint32
		CurHP       uint32
		MaxHP       uint32
		Players     int32
	}
	LastSeen      int64
	LastFound     int64
	LastKilled    int64
	LastUntouched int64 // ?:
}

type FateRelayStateMessage struct {
	A struct {
		WorldID     uint32
		ZoneID      uint32
		InstanceID  uint32
		Coords      Vec3
		ID          uint32
		ReleaseMode ReleaseMode
		Progress    uint8      // 6 - this should be [1-100]
		Status      FateStatus // 7
		StartTime   int64      // 8
		Duration    int32      // 9
	}
	LastSeen      int64
	LastFound     int64
	LastKilled    int64
	LastUntouched int64 // ?:
}

type FateStatus uint8

const (
	FateStatusUnknown FateStatus = iota
	FateStatusPreparation
	FateStatusRunning
	FateStatusComplete
	FateStatusFailed
)

type MessageList []interface{}

func (m *MessageList) DecodeMsgpack(d *msgpack.Decoder) error {
	n, err := d.DecodeArrayLen()
	if err != nil {
		return err
	}
	v := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		v1 := MsgWrapper{}
		err := d.Decode(&v1)
		if err != nil {
			return err
		}
		v = append(v, v1.V)
	}
	*m = v
	return nil
}
