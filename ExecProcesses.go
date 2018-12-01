package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("the Binary: ", binary)

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	/*
		// mine own test
		pbinary, err := exec.LookPath("python3")
		if err != nil {
			panic(err)
		}

		pargs := []string{"python3", "--version"}
		penv := os.Environ()

		execProgP := syscall.Exec(pbinary, pargs, penv)
		if execProgP != nil {
			panic(execProgP)
		}
	*/
}
