package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func init() {
	flag.Parse()
}

func main() {
	command := flag.Args()
	bin, args := command[0], command[1:]

	cmd := exec.Command(bin, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(output))
}
