package main

import "github.com/lheinrichde/gorum/internal/app/gorum"

func main() {
	if err := gorum.Init(); err != nil {
		panic(err)
	}
}
