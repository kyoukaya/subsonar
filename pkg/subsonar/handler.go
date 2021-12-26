package subsonar

func (c *Client) TimeSyncMessageHandler(m *TimeSyncMessage) error { return nil } // NOOP
func (c *Client) LogMessageHandler(m *LogMessage) error           { return nil } // TODO: Maybe we should log these
func (c *Client) SonarMessageHandler(m *SonarMessage) error       { return nil } // TODO: Maybe we should log these
func (c *Client) PingMessageHandler(m *PingMessage) error         { return nil } // TODO: should probably implement proper client behavior
func (c *Client) SeverReadyMessageHandler(m *ServerReadyMessage) error {
	err := c.send(&SonarVersionMessage{
		Version:         2,                      // const
		SonarNetVersion: "0.4.2.3",              //
		Plugin:          "SonarPlugin 0.4.2.5",  // Assembly.GetExecutingAssembly().GetName().Name + " " + VersionUtils.GetSonarPluginVersion()
		Game:            "2021.12.16.0000.0000", // data.GameData.Repositories?["ffxiv"]?.Version ?? "Unknown";
		DalamudVersion:  "6.2.0.13(264c85b)",    // VersionUtils.GetDalamudVersion() + " (" + VersionUtils.GetDalamudBuild() + ")"
		OS:              "10.0",                 // ?
	})
	if err != nil {
		return err
	}
	err = c.send(&PlayerInfoMessage{
		Name:        "Meme Dealer", // Doraemon songs
		HomeWorldID: 72,            // tonberry, probably
	})
	if err != nil {
		return err
	}
	err = c.send(&PlayerPlaceMessage{
		WorldID:    72,
		ZoneID:     370, // kugane
		InstanceID: 0,   // ?: probably 0 indexed. whatever
	})
	if err != nil {
		return err
	}
	err = c.send(&SonarConfig{
		LogLevel: LogLevelInfo,
		HuntConfig: HuntConfig{
			Jurisdiction: map[ExpansionPack]map[HuntRank]SonarJurisdiction{
				ARealmRebornExpansionPack: {
					SSHuntRank:       SonarJurisdictionDatacenter,
					SSMinionHuntRank: SonarJurisdictionNone,
					SHuntRank:        SonarJurisdictionDatacenter,
					AHuntRank:        SonarJurisdictionNone,
					BHuntRank:        SonarJurisdictionNone,
				},
				HeavenswardExpansionPack: {
					SSHuntRank:       SonarJurisdictionDatacenter,
					SSMinionHuntRank: SonarJurisdictionNone,
					SHuntRank:        SonarJurisdictionDatacenter,
					AHuntRank:        SonarJurisdictionNone,
					BHuntRank:        SonarJurisdictionNone,
				},
				StormbloodExpansionPack: {
					SSHuntRank:       SonarJurisdictionDatacenter,
					SSMinionHuntRank: SonarJurisdictionNone,
					SHuntRank:        SonarJurisdictionDatacenter,
					AHuntRank:        SonarJurisdictionNone,
					BHuntRank:        SonarJurisdictionNone,
				},
				ShadowbringersExpansionPack: {
					SSHuntRank:       SonarJurisdictionDatacenter,
					SSMinionHuntRank: SonarJurisdictionNone,
					SHuntRank:        SonarJurisdictionDatacenter,
					AHuntRank:        SonarJurisdictionNone,
					BHuntRank:        SonarJurisdictionNone,
				},
				EndwalkerExpansionPack: {
					SSHuntRank:       SonarJurisdictionDatacenter,
					SSMinionHuntRank: SonarJurisdictionNone,
					SHuntRank:        SonarJurisdictionDatacenter,
					AHuntRank:        SonarJurisdictionDatacenter,
					BHuntRank:        SonarJurisdictionNone,
				},
			},
			JurisdictionOverride: map[uint32]SonarJurisdiction{},
		},
		FateConfig: FateConfig{
			DefaultJurisdiction: SonarJurisdictionNone,
			Jurisdiction: map[uint32]SonarJurisdiction{
				1763: SonarJurisdictionDatacenter, // Devout Pilgrims vs. Daivadipa
				1855: SonarJurisdictionDatacenter, // Omicron Recall: Killing Order
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
