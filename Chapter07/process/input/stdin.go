package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	b := bytes.NewBuffer(nil)
	cmd := exec.Command("cat")
	cmd.Stdin = b
	cmd.Stdout = os.Stdout
	fmt.Fprintf(b, "Hello World! I'm using this memory address: %p", b)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	err := cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
