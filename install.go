package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var to_install string

func install() {
	if to_install == "" {
		return
	}

	fmt.Println("Installing package: ", to_install)
	cmd := exec.Command("yay", "-S", to_install)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal("Executing command err: ", err)
	}

	if err := cmd.Process.Release(); err != nil {
		log.Fatal("Error releasing process, ", err)
	}
}

