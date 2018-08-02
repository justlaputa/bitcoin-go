package main

import "log"

type Config struct {
	SenderKeys   KeyPair
	ReceiverKeys KeyPair
}

func LoadConfig() Config {
	log.Printf("//TODO: load config from file config.json")
	return Config{}
}

func SaveConfig(c Config) {
	log.Printf("//TODO: save config to json")
}
