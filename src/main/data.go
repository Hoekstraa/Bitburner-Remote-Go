package main

type Parameters struct {
	Server   string `json:"server,omitempty"`
	Filename string `json:"filename,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Message struct {
	Jsonrpc string     `json:"jsonrpc,omitempty"`
	Method  string     `json:"method,omitempty"`
	Params  Parameters `json:"params,omitempty"`
	Result  string     `json:"result,omitempty"`
	Id      string     `json:"id,omitempty"`
	Error   string     `json:"error,omitempty"`
}

type Configuration struct {
	RootDir            string `json:"rootDir"`
	FileRemovalAllowed bool   `json:"fileRemovalAllowed"`
}
