package main

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
