package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/fsnotify.v1"
)

func main() {
	dirArgs := options.Parse(os.Args[1:])
	dirArgs = filterPaths(dirArgs)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	for _, dir := range dirArgs {
		subfolders := Subfolders(dir)

		for _, sf := range subfolders {
			err := watcher.Add(sf)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	go func() {
		for {
			select {
			case e := <-watcher.Events:
				switch e.Op {
				case fsnotify.Create:
					fmt.Println("Create", e.Name)
				case fsnotify.Write:
					fmt.Println("Write", e.Name)
				case fsnotify.Remove:
					fmt.Println("Remove", e.Name)
				case fsnotify.Rename:
					fmt.Println("Rename", e.Name)
				}

				job := NewJob(options.Get("cmd").value.(string))
				delay := options.Get("delay").value.(string)
				timer, err := strconv.ParseInt(delay, 10, 64)

				if err == nil {
					time.Sleep(time.Duration(timer) * time.Millisecond)
					go job.Run()
				}
			}
		}
	}()

	<-done
}
