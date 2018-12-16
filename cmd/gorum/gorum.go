package main

import (
	"github.com/ltheinrich/gorum/internal/app/gorum"
)

func main() {
	if err := gorum.Init(); err != nil {
		panic(err)
	}
}
