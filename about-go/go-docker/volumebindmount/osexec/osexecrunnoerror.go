package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	_ = cmd.Run()
	log.Printf("Command finished, no error given because we used _")
}
