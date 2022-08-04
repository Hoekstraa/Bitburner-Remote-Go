package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/grafov/bcast"
)

var configuration = Configuration{
	RootDir:            "/tmp/bit",
	FileRemovalAllowed: false,
}

func main() {
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
