package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	if err := watcher.Add("sample_file"); err != nil {
		panic(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			log.Printf("EVENT: %s\n", event.Name)
		}
	}
}
