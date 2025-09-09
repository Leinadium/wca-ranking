package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Server Server `toml:"server"`
		DB     DB     `toml:"db"`
		WCA    WCA    `toml:"wca"`
		Auth   Auth   `toml:"auth"`
	}

	Server struct {
		Host string `toml:"host"`
	}

	DB struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Tables   struct {
			Default string `toml:"default"`
			Dump    string `toml:"dump"`
		}
	}
	WCA struct {
		ExportURL    string `toml:"export_url"`
		ClientId     string `toml:"client_id"`
		ClientSecret string `toml:"client_secret"`

		Endpoints struct {
			Me             string `toml:"me"`
			LatestData     string `toml:"latest_data"`
			OAuthAuthorize string `toml:"ouath_authorize"`
			OAuthToken     string `toml:"oauth_token"`
		} `toml:"endpoints"`

		RedirectURI string `toml:"redirect_uri"`
	}
	Auth struct {
		RegisterTimeout float32 `toml:"register_timeout"`
	}
)

func New(tomlfile string) *Config {
	var config Config
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		log.Fatalf("Could not read toml: %s", err.Error())
	}
	return &config
}
