package subsonar

type LogLevel uint8

const (
	LogLevelsDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

type SonarConfig struct {
	LogLevel   LogLevel   `msgpack:"logLevel"`
	HuntConfig HuntConfig `msgpack:"huntConfig"`
	FateConfig FateConfig `msgpack:"fateConfig"`
}

func (m *SonarConfig) Opcode() uint8 {
	return 18 //nolint:gomnd
}

type HuntConfig struct {
	Jurisdiction         map[ExpansionPack]map[HuntRank]SonarJurisdiction `msgpack:"jurisdiction"`
	JurisdictionOverride map[uint32]SonarJurisdiction                     `msgpack:"jurisdictionOverride"`
}

type FateConfig struct {
	DefaultJurisdiction SonarJurisdiction            `msgpack:"defaultJurisdiction"`
	Jurisdiction        map[uint32]SonarJurisdiction `msgpack:"jurisdiction"`
}

type ExpansionPack uint8

const (
	UnknownExpansionPack ExpansionPack = iota
	ARealmRebornExpansionPack
	HeavenswardExpansionPack
	StormbloodExpansionPack
	ShadowbringersExpansionPack
	EndwalkerExpansionPack
)

type HuntRank uint8

const (
	NoneHuntRank HuntRank = iota
	BHuntRank
	AHuntRank
	SHuntRank
	SSMinionHuntRank
	SSHuntRank
)

type SonarJurisdiction uint8

const (
	SonarJurisdictionDefault SonarJurisdiction = iota
	SonarJurisdictionNone
	SonarJurisdictionInstance
	SonarJurisdictionZone
	SonarJurisdictionWorld
	SonarJurisdictionDatacenter
	SonarJurisdictionRegion
	SonarJurisdictionAll
)
