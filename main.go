// watch a file for changes and run a specific command
package main

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"log"
	_ "net/http/pprof"
	"os/exec"
)

func run(args []string) {
	cmd_name := args[1]

	var cmd_args []string
	if len(args) > 2 {
		cmd_args = args[2:]
	}

	cmd := exec.Command(cmd_name, cmd_args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error executing the command: ", err)
	}

	log.Printf("%s\n", out)
}

func watch(watcher *fsnotify.Watcher, args []string) {
	file := args[0]
	err := watcher.Add(file)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			//log.Printf("event: %v\n", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				//log.Printf("change to file: %v\n", event.Name)
				run(args)
			}
		case err := <-watcher.Errors:
			log.Println("Error from fsnotify watcher: ", err)
		}
	}
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Not enough arguments, need at least 2")
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	go watch(w, args)

	done := make(chan bool)
	<-done
}
