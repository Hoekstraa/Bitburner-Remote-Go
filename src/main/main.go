package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/grafov/bcast"
)

//go:generate go run web/includeWebFiles.go

func main() {
	if configuration.WebEnabled {
		go webServer()
	}

	outMessages := bcast.NewGroup()
	go outMessages.Broadcast(0)
	go websocket(outMessages)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
	}
	defer watcher.Close()
	recursivelyWatch(watcher, configuration.RootDir)
	go fileWatcher(watcher, outMessages)

	var input string
	fmt.Scanln(&input)
}
