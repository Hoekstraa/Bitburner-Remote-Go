package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"strconv"
)

func pushFile(e fsnotify.Event, msgCount int) Message {
	content, err := os.ReadFile(e.Name) // the file is inside the local directory
	if err != nil {
		fmt.Println(err)
		return Message{
			Jsonrpc: "2.0",
		}
	}

	msg := Message{
		Method: "pushFile",
		Params: Parameters{Server: "home", Filename: gameFriendlyLocation(e.Name), Content: string(content)},
		Id:     strconv.Itoa(msgCount),
	}
	return msg
}

func deleteFile(e fsnotify.Event, msgCount int) Message {
	msg := Message{
		Method: "deleteFile",
		Params: Parameters{Server: "home", Filename: gameFriendlyLocation(e.Name)},
		Id:     strconv.Itoa(msgCount),
	}
	return msg
}
