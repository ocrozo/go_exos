package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/howeyc/fsnotify"
	"github.com/peterbourgon/diskv"
)

func main() {
	fmt.Println("File watch server starting...")
	/**************************************************************/
	d := diskv.New(diskv.Options{
		BasePath:     "files",
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})
	//******** Read file list ***********************************/
	folder := "/tmp/files"
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in " + folder)
	for _, file := range files {
		fmt.Println(file.Name())
		data, _ := ioutil.ReadFile(folder + "/" + file.Name())
		if errdv := d.Write(file.Name(), data); errdv != nil {
			panic(errdv)
		}
	}
	//************* File watcher init *******************/
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case event := <-watcher.Event:
				log.Println("event:", event)
				key := filepath.Base(event.Name)
				log.Println("Modified file: " + key)
				if event.IsDelete() {
					if errdv := d.Erase(key); errdv != nil {
						panic(errdv)
					}
				}
				if event.IsModify() {
					data, _ := ioutil.ReadFile(folder + "/" + key)
					if errdv := d.Write(key, data); errdv != nil {
						panic(errdv)
					}
				}
				if event.IsCreate() {
					data, _ := ioutil.ReadFile(folder + "/" + key)
					log.Println("File was read !")
					if errdv := d.Write(key, data); errdv != nil {
						panic(errdv)
					}
				}
				if event.IsRename() {
					if errdv := d.Erase(key); errdv != nil {
						panic(errdv)
					}
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/tmp/files")
	if err != nil {
		log.Fatal(err)
	}
	/********************++WEB*************************************/
	//http.Handle("/foo", fooHandler)

	//**** Handle send of files
	mux := http.NewServeMux()
	//mux.Handle("/get/", apiHandler{})
	mux.HandleFunc("/get/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		file2serve := req.URL.Path[5:]
		fmt.Println("File to serve: " + file2serve)
		////************* search for the file
		value, errdv := d.Read(file2serve)
		if errdv != nil {
			fmt.Println("Not here! : " + req.URL.Path)
			http.NotFound(w, req)
			return
		}

		w.Write(value)
		return

	})

	mux.HandleFunc("/_files", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintln(w, "<html>")
		fmt.Fprintln(w, "<body>")
		//******************************
		var keyCount int
		for key := range d.Keys(nil) {
			//val, err := d.Read(key)
			if err != nil {
				panic(fmt.Sprintf("key %s had no value", key))
			}

			fmt.Fprintln(w, "<a href='/get/"+key+"'>"+key+"</a><br/>")
			fmt.Printf("Key found: %s\n", key)
			keyCount++
		}
		fmt.Fprintln(w, "</body>")
		fmt.Fprintln(w, "</html>")
		fmt.Printf("%d total keys\n", keyCount)
		//************************
	})
	log.Fatal(http.ListenAndServe(":8080", mux))

	//******************** Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()

}
