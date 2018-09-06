package global

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	homedir "github.com/mitchellh/go-homedir"
)

type Config struct {
	CurrentProfile string             `toml:"current_profile"`
	Profiles       map[string]Profile `toml:"profiles"`
}

type Profile struct {
	Name    string `toml:"name"`
	Server  string `toml:"server"`
	Cluster string `toml:"cluster"`
	Token   string `toml:"token"`
}

func GetCurrentProfile() (*Profile, error) {
	homePath, _ := homedir.Dir()
	configPath := fmt.Sprintf("%s/%s/%s", homePath, ".pfi", "config")
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		contains := strings.Contains(err.Error(), "no such file or directory")
		if contains {
			return SetDefaultCurrentProfile()
		} else {
			fmt.Println("Error when reading configuration file")
			return nil, err
		}
	}

	var config Config
	_, err = toml.Decode(string(b), &config)
	if err != nil {
		fmt.Println("Error when reading configuration file")
		return nil, err
	}
	profile := config.Profiles[config.CurrentProfile]

	return &profile, nil
}

func SetDefaultCurrentProfile() (*Profile, error) {
	profile := &Profile{
		Name:    "default",
		Server:  "http://127.0.0.1:3000",
		Cluster: "default",
		Token:   "",
	}
	config := &Config{
		CurrentProfile: "default",
		Profiles:       map[string]Profile{"default": *profile},
	}

	homePath, _ := homedir.Dir()
	pfiPath := fmt.Sprintf("%s/%s", homePath, ".pfi")
	os.MkdirAll(pfiPath, os.ModePerm)

	configPath := fmt.Sprintf("%s/%s", pfiPath, "config")
	f, err := os.Create(configPath)
	if err != nil {
		fmt.Println("Error when creating default configuration file")
		return nil, err
	}

	if err := toml.NewEncoder(f).Encode(config); err != nil {
		fmt.Println("Error when writing default configuration file")
		return nil, err
	}

	return profile, nil
}
