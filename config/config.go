package config

import (
	"encoding/json"
	"os"
)

type BufferConfig struct {
	HandshakeAddress int `json:"handshake_address"`
	PlayerName       int `json:"player_name"`
	ChatMessage      int `json:"chat_message"`
}

type Config struct {
	ListenAddress string       `json:"listen_address"`
	Logs          bool         `json:"logs"`
	Compression   bool         `json:"enable_compression"`
	Threshold     int          `json:"compression_threshold"`
	BufferConfig  BufferConfig `json:"buffer_config"`
}

var (
	TyphoonConfig Config
)

func InitConfig() (err error) {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &TyphoonConfig); err != nil {
		panic(err)
	}
	return
}
