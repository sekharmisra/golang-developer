package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	/* bs := make([]byte, 99999)

	resp.Body.Read(bs)
	fmt.Println("Response:", string(bs)) */

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println("My Custom Inteface implementation")
	fmt.Println("Length of byte", len("bs"))
	//fmt.Println(bs)
	fmt.Println(string(bs))
	return len(bs), nil
}
