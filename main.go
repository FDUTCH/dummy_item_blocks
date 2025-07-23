package main

import (
	"fmt"
	"github.com/FDUTCH/dummy_item_blocks/dummy"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"log/slog"
	"os"
)

func main() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})
	dummy.EnabledLogging = true
	dummy.Register()
	conf, err := readConfig(slog.Default())
	if err != nil {
		panic(err)
	}

	srv := conf.New()

	srv.CloseOnProgramEnd()
	srv.Listen()

	for p := range srv.Accept() {
		_ = p
	}
}

func readConfig(log *slog.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
	} else {
		data, err := os.ReadFile("config.toml")
		if err != nil {
			return zero, fmt.Errorf("read config: %v", err)
		}
		if err := toml.Unmarshal(data, &c); err != nil {
			return zero, fmt.Errorf("decode config: %v", err)
		}
	}
	cfg, err := c.Config(log)
	return cfg, err
}
