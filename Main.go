package main

import (
	"fmt"
	"net/http"
)

func doMagic(writter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writter, "Hello World!") // send data to client
	return
}

func main() {
	http.HandleFunc("/stock", doMagic)  // map function to address
	err := http.ListenAndServe(":80", nil) // listen to http port
	if err != nil {
		fmt.Sprintf("An Error occured: %s\n", err) // Housten we have a problem
	}
}