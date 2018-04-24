package main

import (
	"fmt"
	"net/http"
	"os"
)

func doMagic(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello World!") // send data to client
	return
}

func main() {
	// get port from the environment variable (heroku sets this)
	var port string = os.Getenv("PORT");
	// or get an default value (e.g. for developing)
	if port == "" {
		port = "1337";
	}
	println("using port: " + port);

	http.HandleFunc("/stock", doMagic)  // map function to address
	err := http.ListenAndServe(":" + port, nil) // listen to http port not allowed
	//err := http.ListenAndServe(":1337", nil) // listen to http port not allowed
	if err != nil {
		fmt.Sprintf("An Error occured: %s\n", err) // Housten we have a problem
	}
}