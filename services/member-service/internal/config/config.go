package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MongoConfig struct {
	URI string `json:"uri" xml:"uri" yaml:"uri"`
}

type ServerConfig struct {
	Port  int  `json:"port" xml:"port" yaml:"port"`
	Debug bool `json:"debug" xml:"debug" yaml:"debug"`
}
type RedisConfig struct {
	URI string `json:"uri" xml:"uri" yaml:"uri"`
}

// Config 包含整个配置文件的数据结构
type Config struct {
	Mongo  MongoConfig  `json:"mongo"`
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
}

var AppConfig Config

// LoadConfig 从 JSON 配置文件加载配置
func LoadConfig() error {
	file, err := os.Open("../internal/config/config.json")

	if err != nil {
		return fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return fmt.Errorf("error decoding config file: %v", err)
	}

	log.Println("Configuration loaded successfully.")
	return nil
}
