package main

import (
	"log"
	"os/exec"
)

func main() {
	log.Println("exec command")
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	log.Printf("current time : %s",out)
}
