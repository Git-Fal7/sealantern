package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListenAddress    string            `json:"listen_address"`
	Logs             bool              `json:"logs"`
	Compression      bool              `json:"enable_compression"`
	Threshold        int               `json:"compression_threshold"`
	CompressionLevel int               `json:"compression_level"`
	InfoFowarding    InfoFowardingType `json:"info_fowarding"`
}

type InfoFowardingType string

const (
	InfoFowardingOfflineMode InfoFowardingType = "offline"
	InfoFowardingBungeeMode  InfoFowardingType = "bungee"
)

var (
	LanternConfig Config
)

func InitConfig() (err error) {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &LanternConfig); err != nil {
		panic(err)
	}
	return
}
