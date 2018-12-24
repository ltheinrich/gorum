package main

import (
	"github.com/ltheinrich/gorum/cmd"
)

func main() {
	if err := cmd.Init(); err != nil {
		panic(err)
	}
}
