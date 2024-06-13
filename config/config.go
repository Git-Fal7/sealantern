package config

import (
	"encoding/base64"
	"encoding/json"
	"image"
	_ "image/png"
	"log"
	"os"
)

type BufferConfig struct {
	HandshakeAddress int `json:"handshake_address"`
	PlayerName       int `json:"player_name"`
	ChatMessage      int `json:"chat_message"`
}

type Config struct {
	ListenAddress string       `json:"listen_address"`
	MaxPlayers    int          `json:"max_players"`
	Motd          string       `json:"motd"`
	Restricted    bool         `json:"restricted"`
	Logs          bool         `json:"logs"`
	Compression   bool         `json:"enable_compression"`
	Threshold     int          `json:"compression_threshold"`
	BufferConfig  BufferConfig `json:"buffer_config"`
	Favicon       string
}

var (
	TyphoonConfig Config
)

func InitConfig() (err error) {
	imgFile, _ := os.Open("./favicon.png")
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)

	if err == nil {
		b := img.Bounds()

		if err == nil {
			if b.Max.X != 64 || b.Max.Y != 64 {
				log.Printf("Invalid icon for server (resize to 64x64 pixels?)")
			}

			fav, err := os.ReadFile("./favicon.png")
			if err == nil {
				TyphoonConfig.Favicon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(fav)
			}
		}
	}

	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &TyphoonConfig); err != nil {
		panic(err)
	}
	return
}
