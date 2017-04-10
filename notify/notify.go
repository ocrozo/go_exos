package main

import (
	"fmt"
	"log"

	"github.com/howeyc/fsnotify"
	//io/ioutil
	//net/http
)

func main() {
	fmt.Println("Hola mundo")
	/**************************************************************/
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				log.Println("File name: ", ev.Name)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/tmp/files")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}
