package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// first command test.
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date ")
	fmt.Println(dateOut)
	fmt.Println(string(dateOut))

	// second test command.
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// .Start() is non pausing.
	grepCmd.Start()

	// .Write is what you hand to to the grep command.
	grepIn.Write([]byte("hello grep\nhellow people\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)

	// runs command but waits.
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// third and last test command.
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(lsCmd.Args)
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
