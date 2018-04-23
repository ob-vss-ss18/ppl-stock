package main

import (
	"fmt"
	"net/http"
)

func doMagic(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello World!") // send data to client
	return
}

func main() {
	http.HandleFunc("/stock", doMagic)  // map function to address
	err := http.ListenAndServe(":1337", nil) // listen to http port not allowed
	if err != nil {
		fmt.Sprintf("An Error occured: %s\n", err) // Housten we have a problem
	}
}