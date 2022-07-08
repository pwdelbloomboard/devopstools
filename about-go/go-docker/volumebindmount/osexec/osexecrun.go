package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleepst", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	if err != nil {
		panic("this is a panic that is never supposed to happen!")
	} else {
		log.Printf("Command finished with error: %v", err)
	}
}
