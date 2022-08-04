package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
)

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

var rootDir = "/tmp/bit"
var fileRemovalAllowed = true

func main() {
	c := make(chan Message)
	go websocket(c)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
	}
	defer watcher.Close()
	recursivelyWatch(watcher, rootDir)
	go fileWatcher(watcher, c)

	var input string
	fmt.Scanln(&input)
}
