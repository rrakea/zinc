package main

import (
	"log"
	"os"
	"os/exec"
)

var to_install string

func install() {
	if to_install == "" {
		return
	}

	cmd := exec.Command("yay", "-S",  "--needed", to_install)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal("Executing command err: ", err)
	}
}
