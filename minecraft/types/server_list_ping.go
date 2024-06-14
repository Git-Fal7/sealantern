package types

type ServerListPing struct {
	Version     ServerListPingVersion     `json:"version"`
	Players     ServerListPingPlayers     `json:"players"`
	Description ServerListPingDescription `json:"description"`
	Favicon     string                    `json:"favicon"`
	ModInfo     ServerListPingModInfo     `json:"modinfo"`
}

type ServerListPingVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type ServerListPingPlayers struct {
	MaxPlayers    int                    `json:"max"`
	OnlinePlayers int                    `json:"online"`
	Sample        []ServerListPingSample `json:"sample"`
}

type ServerListPingSample struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

type ServerListPingDescription struct {
	Motd string `json:"text"`
}

type ServerListPingModInfo struct {
	ModType string   `json:"type"`
	ModList []string `json:"modList"`
}
