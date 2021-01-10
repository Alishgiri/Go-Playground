package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func httpExample() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// USING empty byteSlice to capture RESPONSE data
	// bs := make([]byte, 99999)
	// n, readErr := res.Body.Read(bs)
	// if readErr != nil {
	// 	fmt.Println("Read Error: ", readErr)
	// }
	// fmt.Println(string(bs))
	// fmt.Println(n)

	// USING io to capture RESPONSE data
	// io.Copy(os.Stdout, res.Body)

	// USING our own CUSTOM WRITE function
	lw := logWriter{}
	io.Copy(lw, res.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
