// Package instaweb is a command to very quickly serve a file.  Nothing more.
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		abort(errors.New("Too few arguments."))
	}
	port := os.Args[1]
	filename := os.Args[2]
	validatePort(port)
	validateFilename(filename)
	serve(port, filename)
}

// abort displays the command usage, and exits.
func abort(err error) {
	fmt.Println(err)
	fmt.Println("Usage: instaweb port afile.html")
	os.Exit(0)
}

// serve starts serving the specified file on the specified port.
func serve(port string, filename string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on", port, "; serving", filename, "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		abort(err)
		os.Exit(1)
	}
}

// validatePort ensures the specified port is numeric.
func validatePort(port string) {
	// Ensure port is number.
	_, err := strconv.Atoi(port)
	if err != nil {
		abort(err)
		os.Exit(1)
	}
}

// validateFilename ensures the specified file exists.
func validateFilename(filename string) {
	// Ensure file exists.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		abort(err)
		os.Exit(1)
	}
}
