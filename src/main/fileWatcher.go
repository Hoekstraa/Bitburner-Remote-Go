package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/grafov/bcast"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func recursivelyWatch(watcher *fsnotify.Watcher, directory string) {
	if elementExists(watcher.WatchList(), directory) {
		return
	}
	err := watcher.Add(directory)
	if err != nil {
		return
	}
	files, _ := os.ReadDir(directory)

	for _, f := range files {
		if f.IsDir() {
			recursivelyWatch(watcher, directory+string(filepath.Separator)+f.Name())
		}
	}
}

func gameFriendlyLocation(filename string) string {
	return strings.Replace(filename, configuration.RootDir+string(filepath.Separator), "", 1)
}

func elementExists(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func fileWatcher(w *fsnotify.Watcher, outMessages *bcast.Group) {
	msgCount := 0
	for {
		select {
		case e, ok := <-w.Events:
			if !ok {
				return
			}
			msgCount++

			if e.Op == fsnotify.Create {
				info, err := os.Stat(e.Name)
				if os.IsNotExist(err) {
					return
				}
				if info.IsDir() {
					err := w.Add(e.Name)
					if err != nil {
						return
					}
				}
			}

			if e.Op == fsnotify.Write {
				msg, err := pushFile(e, msgCount)
				if err == nil {
					outMessages.Send(msg)
				}
			}

			if e.Op == fsnotify.Remove {
				if elementExists(w.WatchList(), e.Name) {
					err := w.Remove(e.Name)
					if err != nil {
						return
					}
				} else if configuration.FileRemovalAllowed {
					outMessages.Send(deleteFile(e, msgCount))
				}
			}

			if e.Op == fsnotify.Rename {
				log.Println("Renaming maybe implemented at a later point")
			}

		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Printf("ERROR: %s", err)
		}
	}
}
