package subsonar

import "github.com/vmihailenco/msgpack"

type PlayerInfoMessage struct {
	Name        string
	HomeWorldID uint32
}

// Treat struct as array during msgpack encode
func (m *PlayerInfoMessage) EncodeMsgpack(e *msgpack.Encoder) error {
	e.StructAsArray(true)
	type tmp PlayerInfoMessage
	if err := e.Encode((*tmp)(m)); err != nil {
		return err
	}
	e.StructAsArray(false)
	return nil
}

func (m *PlayerInfoMessage) Opcode() uint8 {
	return 16 //nolint:gomnd
}

type SonarVersionMessage struct {
	Version         int    `msgpack:"version"`
	SonarNetVersion string `msgpack:"sonarNETVersion"`
	Plugin          string `msgpack:"plugin"`
	Game            string `msgpack:"game"`
	DalamudVersion  string `msgpack:"dalamudVersion"`
	OS              string `msgpack:"os"`
}

func (m *SonarVersionMessage) Opcode() uint8 {
	return 19 //nolint:gomnd
}

type PlayerPlaceMessage struct {
	WorldID    uint32
	ZoneID     uint32
	InstanceID uint32
}

func (m *PlayerPlaceMessage) Opcode() uint8 {
	return 17 //nolint:gomnd
}

func (m *PlayerPlaceMessage) EncodeMsgpack(e *msgpack.Encoder) error {
	e.StructAsArray(true)
	type tmp PlayerPlaceMessage
	if err := e.Encode((*tmp)(m)); err != nil {
		return err
	}
	e.StructAsArray(false)
	return nil
}
