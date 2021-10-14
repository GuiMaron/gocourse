package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main () {

	args   	:= os.Args

	if len(args) < 2 {
		log.Fatalln("Missing argument: filename to open")
	}

	lastArg	:= args[len(args) - 1:][0]
	file, e	:= os.Open(lastArg)

	if e != nil {
		log.Fatalln("Impossible to read file", lastArg)
	}

	io.Copy(os.Stdout, file)
	fmt.Println("")

}