package global

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pathfinder-cm/pathfinder-go-client/ext"
)

type Config struct {
	CurrentProfile string `toml:"current_profile"`
	Profiles       map[string]Profile
}

type Profile struct {
	Name    string
	Server  string
	Cluster string
	Token   string
}

func GetCurrentProfile() Profile {
	b, err := ioutil.ReadFile("config.toml")
	if err != nil {
		fmt.Println("Error when reading configuration file")
	}

	var config Config
	_, err = toml.Decode(string(b), &config)
	if err != nil {
		fmt.Println("Error when reading configuration file")
	}

	return config.Profiles[config.CurrentProfile]
}

func NewClient() ext.Client {
	p := GetCurrentProfile()
	h := &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 60 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
		},
	}
	paths := map[string]string{
		"GetNodes":      "api/v1/ext_app/nodes",
		"GetNode":       "api/v1/ext_app/nodes",
		"GetContainers": "api/v1/ext_app/containers",
		"GetContainer":  "api/v1/ext_app/container",
	}
	return ext.NewClient(p.Cluster, p.Token, h, p.Server, paths)
}
