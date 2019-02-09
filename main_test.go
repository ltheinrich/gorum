package main

import (
	"testing"
	"time"

	"github.com/ltheinrich/gorum/cmd"
)

func TestMain(t *testing.T) {
	// make channel for boolean
	fail := make(chan bool)

	// gofunction to call main
	go func() {
		// call main
		main()

		// main ran through, send fail
		fail <- true
	}()

	// gofunction for timeout
	go func() {
		// let main function be running for 3 seconds
		time.Sleep(3 * time.Second)

		// main function did not finish
		fail <- false
	}()

	// check whether failed
	if <-fail {
		// print error message
		t.Error("main function ran through")
	} else {
		// close server
		cmd.Server.Close()
	}
}
