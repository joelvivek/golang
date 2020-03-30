package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("response", resp)

	// bs := make([]byte, 99999) // Give empty space for 99999 elements
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// The above commentd code is equivalent to 
	io.Copy(os.Stdout, resp.Body)

}
