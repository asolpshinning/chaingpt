package gpt

import (
	"log"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

// This function watches a git folder for changes and automatically commits and pushes the changes to the Git folder repository.
func GitAutoCommitPush(folderPath string, timeInterval int) {

	// Watch for changes to the project folder... let us see if this works
	watcher, err := fsnotify.NewWatcher() // so made some changes
	if err != nil {
		log.Fatal(err)
	}
	//defer watcher.Close()

	err = watcher.Add(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Trigger Git commands when a change is detected
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				// Git commands to add, commit, and push changes
				cmdAdd := exec.Command("git", "add", ".")
				cmdAdd.Dir = folderPath

				cmdCommit := exec.Command("git", "commit", "-m", "Automatic commit at "+time.Now().String()+"")
				cmdCommit.Dir = folderPath

				cmdPush := exec.Command("git", "push")
				cmdPush.Dir = folderPath

				// Run the Git commands in sequence
				if err := cmdAdd.Run(); err != nil {
					log.Fatal(err)
				}

				if err := cmdCommit.Run(); err != nil {
					log.Fatal(err)
				}

				if err := cmdPush.Run(); err != nil {
					log.Fatal(err)
				}

				log.Println("\n Changes committed and pushed.")
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
		time.Sleep(time.Duration(timeInterval) * time.Second) // Sleep for 15 seconds before checking again
	}
}
