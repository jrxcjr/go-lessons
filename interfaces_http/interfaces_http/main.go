package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	//syntax for making an http call
	resp, err := http.Get("https://google.com")

	//Error handling must come before treating the response, unlike js.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	//The io copy func, implements all of the functions of the read and write interfaces.
	//It takes two params, one that implements the read interface and another that implements the write interface.
	io.Copy(lw, resp.Body)

	// //another syntax to make the same stuff below

	// //a byte slice is initialized empty with 99999 spaces
	// bs := make([]byte, 99999)
	// //The body pointer is added to the byte slice
	// resp.Body.Read(bs)
	// //the body is printed after being converted from a pointer to a value to a string
	// fmt.Println(string(bs))

	// //The body can be read by using the io.readall function
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// bs := string(body)
	// fmt.Println(bs)
}

// This is the function that satisfies the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	return fmt.Println(string(bs))
}
