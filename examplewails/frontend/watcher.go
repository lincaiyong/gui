package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lincaiyong/log"
	"os"
	"os/exec"
	"path/filepath"
)

func watchFile(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.FatalLog("fail to create watcher: %v", err)
	}
	defer func() { _ = watcher.Close() }()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					cmd := exec.Command("go", "run", "frontend.go")
					log.InfoLog("go run frontend.go")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err = cmd.Run(); err != nil {
						log.ErrorLog("fail to go run: %v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.ErrorLog("error watching file: %v", err)
			}
		}
	}()

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.FatalLog("fail to get abs path: %v", err)
	}
	err = watcher.Add(absPath)
	if err != nil {
		log.FatalLog("fail to watch: %v", err)
	}
	<-done
}

func main() {
	watchFile("frontend.go")
}
