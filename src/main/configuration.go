package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	FileRemovalAllowed bool   `json:"fileRemovalAllowed"`
	RFAPort            int    `json:"rfaPort"`
	RootDir            string `json:"rootDir"`
	WebEnabled         bool   `json:"webEnabled"`
	WebPort            int    `json:"webPort"`
}

var configuration = Configuration{
	FileRemovalAllowed: false,
	RFAPort:            12525,
	RootDir:            ".",
	WebEnabled:         true,
	WebPort:            12526,
}

func loadConfiguration() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Unable to find a User Configuration Directory")
		return
	}
	configDir = configDir + string(os.PathSeparator) + "bitburner-remote"

	_, err = os.ReadDir(configDir)
	if err != nil {
		err = os.Mkdir(configDir, 0750)
		if err != nil {
			fmt.Println("Configuration directory couldn't be made")
			return
		}
	}
	data, err := os.ReadFile(configDir + string(os.PathSeparator) + "config.json")
	if err != nil {
		fmt.Println("Configuration file wasn't found, making one...")
		saveConfiguration()
		return
	}
	conf := Configuration{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("Configuration file seems to contain invalid data, continuing with defaults.")
		return
	}
	configuration = conf
}

func saveConfiguration() {
	conf, err := json.Marshal(configuration)
	if err != nil {
		fmt.Println("Failed to write configuration")
		return
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Unable to find a User Configuration Directory")
		return
	}
	filename := configDir + string(os.PathSeparator) + "bitburner-remote" + string(os.PathSeparator) + "config.json"

	_, err = os.Lstat(filename)
	if err != nil {
		_, err = os.Create(filename)
		if err != nil {
			fmt.Println("Failed to create configuration file")
			return
		}
	}
	err = os.WriteFile(filename, conf, 0750)
	if err != nil {
		fmt.Println("Failed to write configuration file")
		return
	}
}
