package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func (logWriter) Write (data []byte) (int, error) {
	return os.Stdout.Write(data)
}

func main () {

	resp, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	data, readErr := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if readErr != nil {
		fmt.Println("Error: ", readErr)
		os.Exit(1)
	}

	//fmt.Println(string(data))

	lw := logWriter{}
	lw.Write(data)

	io.Copy(lw, resp.Body)

}
